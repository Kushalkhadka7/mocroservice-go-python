import os
import json

from airflow.constants import constants

from airflow.models import DAG
from airflow.utils.dates import days_ago
from airflow.contrib.hooks.fs_hook import FSHook
from airflow.operators.python_operator import PythonOperator
from airflow.sensors.external_task_sensor import ExternalTaskSensor


args = {
    'owner': constants.OWNER,
    'start_date': days_ago(1)
}

dag = DAG(
    default_args=args,
    dag_id="my_sample_dag_2",
    schedule_interval=constants.DAG_RUN_TIME_INTRVAL
)


def write_file(**context):
    hook = FSHook('fs_default')
    base_path = hook.get_path()
    file_path = os.path.join(base_path, constants.JSON_FILE_NAME)

    with open(file_path, 'r') as json_file:
        loaded_file = json.load(json_file)

        print(loaded_file["0"])

    print(json_file)


with dag:
    external_sensor = ExternalTaskSensor(
        task_id='external_sensor',
        external_dag_id='my_sample-dag',
        external_task_id='read_file'
    )

    write_file = PythonOperator(
        task_id="write_file",
        python_callable=write_file,
        provide_context=True
    )

    external_sensor >> write_file
