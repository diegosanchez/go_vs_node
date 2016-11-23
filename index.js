// Retrieve
var MongoClient = require('mongodb').MongoClient;
var http = require('http');

// Connect to the db
MongoClient.connect("mongodb://127.0.0.1:27017/local", function(err, db) {
    if(!err) {
        console.log("We are connected");
    }

    var server = http.createServer( function handleRequest(req, res){
        db.collection("startup_log", function(err, coll) {
            if ( err ) {
                process.exit(1);
            }

            coll.findOne({}, function(err, entry) {
                res.end(JSON.stringify(entry));
            });
        });
    });
    server.listen(9999, function(){
        console.log("Server listening on: http://localhost:%s", 9999 );
    });
});
