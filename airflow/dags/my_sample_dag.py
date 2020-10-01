from airflow.models import DAG
from airflow.utils.dates import days_ago
from airflow.contrib.sensors.file_sensor import FileSensor
from airflow.operators.python_operator import PythonOperator

import random

args={
    'owner':'kushal khadka',
    'start_date': days_ago(1)
}

dag = DAG(
    dag_id = 'my_sample-dag',
    default_args = args,
    schedule_interval = None
)

def task1_func(**context):
    randoInt = random.random()
    print(randoInt)
    context['ti'].xcom_push(key='randomInt',value=randoInt)
    print(context)
    print("hello i am task1_func")

def task2_func(**context):
    received_value = context['ti'].xcom_pull(key='randomInt')
    print(received_value)
    print("hello i am task2_func")

def task3_func(**context):
    print("hello i am task3_func and i am run after file sensing task is completed.")

with dag:
    sensing_task = FileSensor(
        poke_interval = 10,
        filepath='text.txt',
        task_id='sensing-task',
        fs_conn_id = 'fs_default'
    )

    task1 = PythonOperator(
        task_id='task1', 
        python_callable = task1_func,
        provide_context=True
    )

    task2 = PythonOperator(
        task_id='task2',
        python_callable = task2_func,
        provide_context=True
    )

    task3 = PythonOperator(
        task_id='task3',
        python_callable = task3_func,
        provide_context=True
    )

    sensing_task >> task3