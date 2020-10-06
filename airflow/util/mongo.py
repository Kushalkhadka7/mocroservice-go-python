from pymongo import MongoClient


class MongoDB:
    """
    Initializes connection to mongodb database with 
    the provided port, host and databse name.
    Establish connection the the required collection.
    Return the collection ref.
    """

    def __init__(self, port, host, db):
        self.db = db
        self.port = port
        self.host = host

    def get_db_conn(self):
        return MongoClient(self.host, self.port)

    def get_monog_client(self):
        return self.get_db_conn()[self.db]

    def get_collection_ref(self, collection):
        return self.get_monog_client()[collection]
