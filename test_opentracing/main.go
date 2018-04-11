package main

import (
	"fmt"
	"time"
	"bytes"
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	zipkin "github.com/openzipkin/zipkin-go-opentracing"
	"github.com/opentracing/opentracing-go/log"
	"errors"
)

var tracer opentracing.Tracer

func initTracer(args... string) {
	if len(args) < 1 {
		panic("Init args length < 1")
	}

	addr := args[0]
	debugStr := args[1]
	hostPort := args[2]
	serviceName := args[3]

	debug := false

	if debugStr == "true" {
		debug = true
	}

	initZipkin(addr, debug, hostPort, serviceName)

	opentracing.InitGlobalTracer(tracer)
}

func initZipkin(addr string, debug bool, hostPort string, serviceName string) {
	// create collector.
	collector, err := zipkin.NewHTTPCollector(addr)

	if err != nil {
		panic(err)
	}

	// create recorder.
	recorder := zipkin.NewRecorder(collector, debug, hostPort, serviceName)

	// create tracer.
	tracer, err = zipkin.NewTracer(
		recorder,
		zipkin.ClientServerSameSpan(true),
		zipkin.TraceID128Bit(true),
	)

	if err != nil {
		panic(err)
	}
}

func testClient(ctx context.Context) {
	//Create root span using rpc kind
	opts := []opentracing.StartSpanOption{
		ext.SpanKindRPCClient,
	}

	clientSpan := tracer.StartSpan("test_opentracing_client", opts...)
	defer clientSpan.Finish()

	time.Sleep(time.Millisecond * 10)

	//Inject
	tracingBufClient := bytes.NewBuffer([]byte{})
	if err := tracer.Inject(clientSpan.Context(), opentracing.Binary, tracingBufClient); err != nil {
		fmt.Printf("grpc_opentracing: failed serializing trace information: %v\r\n", err)
	}

	testServer(tracingBufClient.Bytes())
}

func testServer(tracingInfo []byte) {
	//Extract
	tracingBufClient := bytes.NewBuffer(tracingInfo)
	parentSpanContext, err := tracer.Extract(opentracing.Binary, tracingBufClient)
	if err != nil && err != opentracing.ErrSpanContextNotFound {
		fmt.Printf("grpc_opentracing: failed parsing trace information: %v\r\n", err)
	}

	opts := []opentracing.StartSpanOption{
		ext.RPCServerOption(parentSpanContext),
		ext.SpanKindRPCServer,
	}

	serverSpan := tracer.StartSpan("test_opentracing_server", opts...)
	defer serverSpan.Finish()

	time.Sleep(time.Millisecond * 10)

	//Inject
	tracingBufServer := bytes.NewBuffer([]byte{})
	if err := tracer.Inject(serverSpan.Context(), opentracing.Binary, tracingBufServer); err != nil {
		fmt.Printf("grpc_opentracing: failed serializing trace information: %v\r\n", err)
	}

	go testProducer(tracingBufServer.Bytes())
}

func testProducer(tracingInfo []byte) {
	//Extract
	tracingBufServer := bytes.NewBuffer(tracingInfo)
	parentSpanContext, err := tracer.Extract(opentracing.Binary, tracingBufServer)
	if err != nil && err != opentracing.ErrSpanContextNotFound {
		fmt.Printf("grpc_opentracing: failed parsing trace information: %v\r\n", err)
	}

	opts := []opentracing.StartSpanOption{
		ext.RPCServerOption(parentSpanContext),
		ext.SpanKindProducer,
	}

	producerSpan := tracer.StartSpan("test_opentracing_producer", opts...)
	defer producerSpan.Finish()

	time.Sleep(time.Millisecond * 1)

	//Inject
	tracingBufProducer := bytes.NewBuffer([]byte{})
	if err := tracer.Inject(producerSpan.Context(), opentracing.Binary, tracingBufProducer); err != nil {
		fmt.Printf("grpc_opentracing: failed serializing trace information: %v\r\n", err)
	}

	go testConsumer(tracingBufProducer.Bytes())
	time.Sleep(time.Millisecond * 5)
}

func testConsumer(tracingInfo []byte) {
	//Extract
	tracingBufProducer := bytes.NewBuffer(tracingInfo)
	parentSpanContext, err := tracer.Extract(opentracing.Binary, tracingBufProducer)
	if err != nil && err != opentracing.ErrSpanContextNotFound {
		fmt.Printf("grpc_opentracing: failed parsing trace information: %v\r\n", err)
	}

	opts := []opentracing.StartSpanOption{
		ext.RPCServerOption(parentSpanContext),
		ext.SpanKindConsumer,
	}

	consumerSpan := tracer.StartSpan("test_opentracing_consumer", opts...)
	defer consumerSpan.Finish()

	time.Sleep(time.Millisecond * 10)

	err = errors.New("test error!!!")
	if err != nil {
		ext.Error.Set(consumerSpan, true)
		consumerSpan.LogFields(log.String("event", "error"), log.String("message", err.Error()))
	}

	time.Sleep(time.Millisecond * 10)

	consumerSpan.LogKV("test_key", "test_value")

	time.Sleep(time.Millisecond * 10)

	consumerSpan.SetTag("tag_key", "tag_value")

	time.Sleep(time.Millisecond * 10)
}

func main() {
	initTracer("http://localhost:9411/api/v1/spans", "false", "localhost:58888", "demo")
	go testClient(context.Background())

	time.Sleep(time.Second * 3)
}
