package v1

import (
	scheme "github.com/openshift/client-go/image/clientset/scheme"
	image_v1 "github.com/openshift/origin/pkg/image/apis/image/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
)

// ImageStreamImagesGetter has a method to return a ImageStreamImageInterface.
// A group's client should implement this interface.
type ImageStreamImagesGetter interface {
	ImageStreamImages(namespace string) ImageStreamImageInterface
}

// ImageStreamImageInterface has methods to work with ImageStreamImage resources.
type ImageStreamImageInterface interface {
	Get(name string, options v1.GetOptions) (*image_v1.ImageStreamImage, error)
	ImageStreamImageExpansion
}

// imageStreamImages implements ImageStreamImageInterface
type imageStreamImages struct {
	client rest.Interface
	ns     string
}

// newImageStreamImages returns a ImageStreamImages
func newImageStreamImages(c *ImageV1Client, namespace string) *imageStreamImages {
	return &imageStreamImages{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the imageStreamImage, and returns the corresponding imageStreamImage object, and an error if there is any.
func (c *imageStreamImages) Get(name string, options v1.GetOptions) (result *image_v1.ImageStreamImage, err error) {
	result = &image_v1.ImageStreamImage{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("imagestreamimages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}
