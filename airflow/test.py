import os
import json

from pymongo import MongoClient

import constants


def load_data_to_monog():
    db = get_db_client()
    collection = db['test']
    json_data = write_to_json_file()
    array_of_data = []

    for json in json_data.items:
        print(json)

    data = collection.insert(json_data['0'])

    print(data)


def write_to_json_file():
    base_path = './files/'
    file_path = os.path.join(base_path, 'laptop.json')

    with open(file_path, 'r') as json_file:
        loaded_file = json.load(json_file)

    return loaded_file


def get_db_client():
    client = MongoClient("localhost", 27017)
    db = client['test']

    return db


def main():
    load_data_to_monog()


if __name__ == '__main__':
    main()
