mongo -- "$MONGO_INITDB_DATABASE" <<EOF
    var admin = db.getSiblingDB('admin');
    admin.auth('$MONGO_INITDB_ROOT_USERNAME', '$MONGO_INITDB_ROOT_PASSWORD');

    use admin
    db.createUser(
      {
        "user": '$MONGO_USERNAME',
        "pwd": '$MONGO_PASSWORD',
        "roles": [
            {
              role: "readWrite",
              db: "admin"
            },
            {
               role: "readWrite",
               db: '$MONGO_DB'
            }
          ],
          "mechanisms" : [
          		"SCRAM-SHA-1",
          		"SCRAM-SHA-256"
          	]
        }
    );
EOF