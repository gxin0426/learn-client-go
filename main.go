package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"strings"
	"time"
)

func main() {
	//生成config
	config, err := clientcmd.BuildConfigFromFlags("", "config")
	if err != nil {
		panic(err)
	}




	//config.APIPath = "api"
	//config.GroupVersion = &corev1.SchemeGroupVersion
	//config.NegotiatedSerializer = scheme.Codecs
	////restclient
	//restClinet, err := rest.RESTClientFor(config)
	//if err != nil {
	//	panic(err)
	//}
	////调用request的方法 输出url  restclient通过 go 的newrequest 发出请求
	//str := restClinet.Get().Namespace("default").Resource("pods").VersionedParams(&metav1.ListOptions{Limit:1}, scheme.ParameterCodec).URL().String()
	//fmt.Println(str)

	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err)
	//}
	//podClient := clientset.CoreV1().Pods("")
	//list, err := podClient.List(metav1.ListOptions{Limit:1})
	//if err != nil {
	//	panic(err)
	//}
	//for _, d := range list.Items {
	//	fmt.Println(d.Namespace, d.Name, d.Annotations)
	//}


	////创建一个list watch client
	//podListWatcher := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "pods", v1.NamespaceDefault, fields.Everything())
	////创建一个限速队列
	//queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	//
	//indexer, informer := cache.NewIndexerInformer(podListWatcher, &v1.Pod{}, 0, cache.ResourceEventHandlerFuncs{
	//	AddFunc: func(obj interface{}) {
	//		key, err := cache.MetaNamespaceKeyFunc(obj)
	//		if err == nil {
	//			queue.Add(key)
	//			fmt.Println("queue len add : ", queue.Len())
	//		}
	//	},
	//	UpdateFunc: func(oldObj, newObj interface{}) {
	//		key, err := cache.MetaNamespaceKeyFunc(newObj)
	//		if err == nil {
	//			queue.Add(key)
	//			fmt.Println("queue len update : ", queue.Len())
	//		}
	//	},
	//	DeleteFunc: func(obj interface{}) {
	//		key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	//		if err == nil {
	//			queue.Add(key)
	//			fmt.Println("queue len delete : ", queue.Len())
	//		}
	//	},
	//}, cache.Indexers{})
	//controller := NewController(queue, indexer, informer)
	//
	//indexer.Add(&v1.Pod{
	//
	//	ObjectMeta: metav1.ObjectMeta{
	//		Name: "etcd-sample-2",
	//		Namespace: v1.NamespaceDefault,
	//	},
	//
	//})
	//fmt.Println("此刻队列的长度 ： ", queue.Len())
	//stop := make(chan struct{})
	//defer close(stop)
	//go controller.Run(10, stop)
	//select {
	//
	//}

	//dynamic client
	dynaClient, err := dynamic.NewForConfig(config)
	gvr := schema.GroupVersionResource{"devops.my.domain", "v1", "etcds"}
	unstrcObj, err := dynaClient.Resource(gvr).List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(unstrcObj)


	//discoveryClient  code
	//dClient, err := discovery.NewDiscoveryClientForConfig(config)
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, apiResouceList, err := dClient.ServerGroupsAndResources()
	//if err != nil {
	//	panic(err)
	//}
	//
	//for _, list := range apiResouceList {
	//	gv, err := schema.ParseGroupVersion(list.GroupVersion)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	for _, reouce := range list.APIResources {
	//		fmt.Println(reouce.Name," :: ", gv.Group,"   :::    ", gv.Version)
	//	}
	//}
	//获取一个pod 的 最新的 resourceversion
	//pc, _ := clientset.CoreV1().Pods("kube-system").Get("kube-proxy-rc5tb", v1.GetOptions{})
	//fmt.Println(pc.ResourceVersion)

	//informer机制的实现
	//stopCh := make(chan struct{})
	//defer close(stopCh)
	//
	//shareInformers := informers.NewSharedInformerFactory(clientset, 0)
	//
	//podInformer := shareInformers.Core().V1().Pods().Informer()
	//
	//podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
	//	AddFunc: func(obj interface{}) {
	//		mObj := obj.(v1.Object)
	//		log.Printf("new pod added to store : %s", mObj.GetName())
	//
	//	},
	//	UpdateFunc: func(oldObj, newObj interface{}) {
	//		oObj := oldObj.(v1.Object)
	//		nObj := newObj.(v1.Object)
	//		log.Printf("%s pod update to %s", oObj.GetName(), nObj.GetName())
	//	},
	//	DeleteFunc: func(obj interface{}) {
	//		mObj := obj.(v1.Object)
	//		log.Printf("pod delete from store : %s", mObj.GetName())
	//
	//	},
	//
	//})
	//
	//podInformer.Run(stopCh)

	//indexer 工作原理
	//index := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{"byUser": UsersIndexFunc})
	//pod1 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "one", Annotations: map[string]string{"users": "ernie,bert"}}}
	//pod2 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "two", Annotations: map[string]string{"users": "bert,oscar"}}}
	//pod3 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "tre", Annotations: map[string]string{"users": "ernie,elmo"}}}
	//
	//index.Add(pod1)
	//index.Add(pod2)
	//index.Add(pod3)
	//
	//erniePods, err := index.ByIndex("byUser", "oscar")
	//if err != nil {
	//	panic(err)
	//}
	////fmt.Println(erniePods)
	//for _, erniePod := range erniePods {
	//	fmt.Println(erniePod.(*v1.Pod).Name)
	//}

}

type Controller struct {
	//indexer 是client-go用来存储资源对象并带索引功能的本地存储 Reflector 从DeltaFIFO中将消费出来的资源对象存储至indexer local store 二级缓存的第二级
	indexer cache.Indexer
	//workqueue 限速队列
	queue workqueue.RateLimitingInterface
	//controller 是部分reflector  list watch 都是在这里实现
	informer cache.Controller
}

func NewController(queue workqueue.RateLimitingInterface, indexer cache.Indexer, informer cache.Controller) *Controller {
	return &Controller{
		indexer:  indexer,
		queue:    queue,
		informer: informer,
	}
}

func (c *Controller) processNextItem() bool {
	//等待直到有一个item进入 working queue
	key, quit := c.queue.Get()
	if quit {
		return false
	}
	//告诉队列我们已经处理完这个key， 为其他的worker去除了障碍， 这样做是线程安全的 因为有相同key的两个pod不能并行的被处理
	defer c.queue.Done(key)
	//调用包含业务逻辑的方法
	err := c.syncToStdout(key.(string))
	c.handlerErr(err, key)
	return true
}
//sync 是控制器的业务逻辑， 控制器简单的打印出关于pod的stdout
func (c *Controller) syncToStdout(key string) error {
	obj, exists, err := c.indexer.GetByKey(key)
	if err  != nil {
		klog.Errorf("获取key 对象 %s from store failed with %v", key, err)
	}

	if !exists {
		fmt.Printf("Pod %s does not exist anymore\n", key)
	}else {
		fmt.Printf("Sync/Add/Update for Pod %s\n", obj.(*v1.Pod).GetName())
	}
	return  nil
}

func (c *Controller)handlerErr(err error, key interface{}) {
	if err == nil {
		c.queue.Forget(key)
		return
	}

	if c.queue.NumRequeues(key) < 5 {
		klog.Infof("Error syncing pod %v: %v", key, err)
		c.queue.AddRateLimited(key)
		return
	}

	c.queue.Forget(key)
	runtime.HandleError(err)
	klog.Infof("Dropping pod %q out of the queue: %v", key, err)
}


func (c *Controller) Run(threadiness int, stopCh chan struct{}) {
	defer runtime.HandleCrash()

	// Let the workers stop when we are done
	defer c.queue.ShutDown()
	klog.Info("Starting Pod controller")

	go c.informer.Run(stopCh)

	// Wait for all involved caches to be synced, before processing items from the queue is started
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh
	klog.Info("Stopping Pod controller")
}

func (c *Controller) runWorker() {
	for c.processNextItem() {
	}
}

func UsersIndexFunc(obj interface{})([]string, error) {
	pod := obj.(*v1.Pod)
	userString := pod.Annotations["users"]
	return strings.Split(userString, ","), nil
}

