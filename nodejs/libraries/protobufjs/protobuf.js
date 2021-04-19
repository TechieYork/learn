import protobuf from 'protobufjs';

protobuf.load('./user.proto', function(err, root) {
    console.log(root.nested);
});

