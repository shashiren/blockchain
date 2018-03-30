import time #时间
import queue #队列

from    multiprocessing.managers import BaseManager #分布式管理器

task_queue=queue.Queue()#任务队列
result_queue=queue.Queue()#结果队列

class QueueMananger(BaseManager):#创建类继承式分布式管理器
    pass

def get_task():
    return task_queue

def get_result():
    return result_queue

if __name__ =="__main__":
    frozenset()

    QueueMananger.register("get_task_queue",callable= get_task)#任务队列
    QueueMananger.register("get_result_queue",callable=get_result)#结果队列

    manager=QueueMananger(address=("127.0.0.1",55000),authkey=b"123456")#服务器密码

    manager.start()#启动
    task_queue=manager.get_task_queue()
    result_queue=manager.get_result_queue()


    print("服务器开始")
    for i in range(100):#加入10个数据
        print("server put data %d" % i)
        task_queue.put(i)

    for i in range(100):#取出10个数据
        print("server get data %s"%str(result_queue.get()) )

    manager.shutdown()#关闭
    print("服务器结束")


