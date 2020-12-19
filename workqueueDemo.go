package main



//
//import (
//	"flag"
//	"fmt"
//	v1 "k8s.io/api/core/v1"
//	"k8s.io/apimachinery/pkg/fields"
//	"k8s.io/apimachinery/pkg/util/runtime"
//	"k8s.io/apimachinery/pkg/util/wait"
//	"k8s.io/client-go/kubernetes"
//	"k8s.io/client-go/tools/cache"
//	"k8s.io/client-go/tools/clientcmd"
//	"k8s.io/client-go/util/workqueue"
//	"k8s.io/klog"
//	"time"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//)
//
//
///*
//此示例演示如何编写（跟踪监视资源状态）的控制器
//
//它演示了如何:
//	1.将工作队列（workqueue）与缓存（cache）合并到完整的控制器
//	2.启动时同步控制器
//这个示例基于： https://git.k8s.io/community/contributors/devel/sig-api-machinery/controllers.md
//
//一个kubernetes controller 是一个主动调谐过程（process），controller观测着系统中某个资源的真实状态，同时也观测着系统中某个资源的期望状态，
//然后controller发出指令 试图使系统某个资源的真实状态 更接近期望状态
//
//*/
//
//
//// Controller demonstrates how to implement a controller with client-go.
////自己实现的controller （用户侧）
//
//
////1.首先我们定义一个Controller结构体
//type Controller struct {
//	// indexer的引用
//	indexer  cache.Indexer
//	//workqueue
//	queue    workqueue.RateLimitingInterface
//	//informer的引用
//	informer cache.Controller
//}
//
//// NewController creates a new Controller.
//func NewController(queue workqueue.RateLimitingInterface, indexer cache.Indexer, informer cache.Controller) *Controller {
//	return &Controller{
//		informer: informer,
//		indexer:  indexer,
//		queue:    queue,
//	}
//}
//
////定义controller工作流
//// Run begins watching and syncing.
//func (c *Controller) Run(threadiness int, stopCh chan struct{}) {
//	defer runtime.HandleCrash()
//
//	// Let the workers stop when we are done
//	defer c.queue.ShutDown()
//	klog.Info("Starting Pod controller")
//	//启动informer
//	go c.informer.Run(stopCh)
//
//	// Wait for all involved caches to be synced, before processing items from the queue is started
//	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
//		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
//		return
//	}
//	//启动多个worker 处理workqueue中的对象
//	for i := 0; i < threadiness; i++ {
//		go wait.Until(c.runWorker, time.Second, stopCh)
//	}
//
//	<-stopCh
//	klog.Info("Stopping Pod controller")
//}
//
//
////从workqueue中取出对象，并打印消息
//func (c *Controller) processNextItem() bool {
//	// Wait until there is a new item in the working queue
//	key, quit := c.queue.Get()
//	if quit {
//		return false
//	}
//	// Tell the queue that we are done with processing this key. This unblocks the key for other workers
//	// This allows safe parallel processing because two pods with the same key are never processed in
//	// parallel.
//	defer c.queue.Done(key)
//
//
//	// Invoke the method containing the business logic
//	//包含了处理逻辑 好像reconcile就在这个里面？
//	// 将key对应的 object 的信息进行打印
//	err := c.syncToStdout(key.(string))
//	// Handle the error if something went wrong during the execution of the business logic
//	c.handleErr(err, key)
//	return true
//}
//
//// syncToStdout is the business logic of the controller. In this controller it simply prints
//// information about the pod to stdout. In case an error happened, it has to simply return the error.
//// The retry logic should not be part of the business logic.
//func (c *Controller) syncToStdout(key string) error {
//	obj, exists, err := c.indexer.GetByKey(key)
//	if err != nil {
//		klog.Errorf("Fetching object with key %s from store failed with %v", key, err)
//		return err
//	}
//
//	if !exists {
//		// Below we will warm up our cache with a Pod, so that we will see a delete for one pod
//		fmt.Printf("Pod %s does not exist anymore\n", key)
//	} else {
//		// Note that you also have to check the uid if you have a local controlled resource, which
//		// is dependent on the actual instance, to detect that a Pod was recreated with the same name
//		fmt.Printf("Sync/Add/Update for Pod %s\n", obj.(*v1.Pod).GetName())
//	}
//	return nil
//}
//
//// handleErr checks if an error happened and makes sure we will retry later.
//func (c *Controller) handleErr(err error, key interface{}) {
//	if err == nil {
//		// Forget about the #AddRateLimited history of the key on every successful synchronization.
//		// This ensures that future processing of updates for this key is not delayed because of
//		// an outdated error history.
//		c.queue.Forget(key)
//		return
//	}
//
//	// This controller retries 5 times if something goes wrong. After that, it stops trying.
//	if c.queue.NumRequeues(key) < 5 {
//		klog.Infof("Error syncing pod %v: %v", key, err)
//
//		// Re-enqueue the key rate limited. Based on the rate limiter on the
//		// queue and the re-enqueue history, the key will be processed later again.
//		c.queue.AddRateLimited(key)
//		return
//	}
//
//	c.queue.Forget(key)
//	// Report to an external entity that, even after several retries, we could not successfully process this key
//	runtime.HandleError(err)
//	klog.Infof("Dropping pod %q out of the queue: %v", key, err)
//}
//
//
///***********具体处理workqueue中对象的流程**************/
//
//func (c *Controller) runWorker() {
//	//启动无线循环 接受并处消息
//	for c.processNextItem() {
//	}
//}
//
//func main() {
//	var kubeconfig string
//	var master string
//
//	flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
//	flag.StringVar(&master, "master", "", "master url")
//	flag.Parse()
//
//	// creates the connection
//	config, err := clientcmd.BuildConfigFromFlags("", "config")
//	if err != nil {
//		klog.Fatal(err)
//	}
//
//	// creates the clientset
//	clientset, err := kubernetes.NewForConfig(config)
//	if err != nil {
//		klog.Fatal(err)
//	}
//	// 指定ListerWatcher 在default ns下监听 pod资源
//	// create the pod watcher
//	podListWatcher := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "pods", v1.NamespaceDefault, fields.Everything())
//
//	// create the workqueue
//	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
//
//	// Bind the workqueue to a cache with the help of an informer. This way we make sure that
//	// whenever the cache is updated, the pod key is added to the workqueue.
//	// Note that when we finally process the item from the workqueue, we might see a newer version
//	// of the Pod than the version which was responsible for triggering the update.
//
//	//client-go tool/cache/controller.go文件 中的controller 也就是informer
//	indexer, informer := cache.NewIndexerInformer(podListWatcher, &v1.Pod{}, 0, cache.ResourceEventHandlerFuncs{
//		//当有pod创建时，根据delta FIFO 弹出的object生成对应的key，并加入到workqueue中。此处可以根据Object的一些属性进行过滤
//		AddFunc: func(obj interface{}) {
//			key, err := cache.MetaNamespaceKeyFunc(obj)
//			if err == nil {
//				queue.Add(key)
//			}
//		},
//		UpdateFunc: func(old interface{}, new interface{}) {
//			key, err := cache.MetaNamespaceKeyFunc(new)
//			if err == nil {
//				queue.Add(key)
//			}
//		},
//		//DeletionHandlingMetaNamespaceKeyFunc 会在生成key之前检查，因为资源删除后有可能会进行重建等操作
//		//监听时错过了删除信息，从而导致该条记录是陈旧的
//		DeleteFunc: func(obj interface{}) {
//			// IndexerInformer uses a delta queue, therefore for deletes we have to use this
//			// key function.
//			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
//			if err == nil {
//				queue.Add(key)
//			}
//		},
//	}, cache.Indexers{})
//
//	controller := NewController(queue, indexer, informer)
//
//	// We can now warm up the cache for initial synchronization.
//	// Let's suppose that we knew about a pod "mypod" on our last run, therefore add it to the cache.
//	// If this pod is not there anymore, the controller will be notified about the removal after the
//	// cache has synchronized.
//	indexer.Add(&v1.Pod{
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      "etcd-sample-0",
//			Namespace: v1.NamespaceDefault,
//		},
//	})
//
//	// Now let's start the controller
//	stop := make(chan struct{})
//	defer close(stop)
//	go controller.Run(1, stop)
//
//	// Wait forever
//	select {}
//}
