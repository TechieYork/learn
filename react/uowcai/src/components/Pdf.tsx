import { Layout, Skeleton, Spin } from "antd";
import { useMemo } from "react";
import { pdfjs, Document, Page } from "react-pdf";
import { useDispatch, useSelector } from "react-redux";
import { setPdfLoading } from "./pdfSlice";

pdfjs.GlobalWorkerOptions.workerSrc = new URL(
  "pdfjs-dist/build/pdf.worker.min.mjs",
  import.meta.url
).toString();

interface PdfProps {
  file: string;
  page: number;
}

function Pdf(props: PdfProps) {
  let file = useMemo(() => {
    return props.file;
  }, [props.file]);

  let loading = useSelector((state: any) => state.pdf.loading);
  let dispatch = useDispatch();
  console.log("page", props.page, "loading", loading, "file", file);

  return (
    <Layout.Content
      style={{
        border: loading ? "" : "1px solid #888888",
        boxShadow: loading ? "" : "2px 2px 3px #888888",
        width: window.innerWidth / 2,
        // minHeight: 480,
        // minWidth: window.innerWidth / 2,
      }}
    >
      <Document
        file={file}
        onLoadSuccess={() => dispatch(setPdfLoading(false))}
        loading={<Skeleton active />}
      >
        <Page
          pageNumber={props.page}
          renderTextLayer={false}
          renderAnnotationLayer={false}
          width={window.innerWidth / 2}
          onLoadSuccess={() => dispatch(setPdfLoading(false))}
          
        />
      </Document>
    </Layout.Content>
  );
}

export { Pdf };
