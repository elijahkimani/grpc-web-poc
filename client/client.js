var grpc = require("@grpc/grpc-js");
var protoLoader = require("@grpc/proto-loader");
var path = require("path");

const PORT = 30043;

var PROTO_PATH = path.resolve(__dirname, "../proto/products.proto");

// transcript the proto (package definition) to JS object
var packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

// load the JS version of the proto to create a GRPC object
var productsProto = grpc.loadPackageDefinition(packageDefinition);
var ProductsService = productsProto.products.ProductsService;

const client = new ProductsService(
  `localhost:${PORT}`,
  grpc.credentials.createInsecure()
);

module.exports = client;
