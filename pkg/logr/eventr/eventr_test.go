package eventr_test

import (
	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	corev1 "k8s.io/api/core/v1"

	"github.com/konflux-ci/project-controller/pkg/logr/eventr"
)

var _ = Describe("Eventr", func() {
	var (
		recorder *record.FakeRecorder
		object runtime.Object
		logger logr.Logger
	)

	BeforeEach(func () {
		recorder = record.NewFakeRecorder(10)
		recorder.IncludeObject = true

		object = &corev1.ConfigMap{}
		Expect(object.GetObjectKind().GroupVersionKind().GroupKind().String()).NotTo(BeEmpty())

		logger = eventr.NewEventr(recorder, object)
	})

	It("Generates events on subject on logging calls", func () {
		logger.Info("Something happened")

		Expect(recorder.Events).Should(Receive(Equal(
			"Something happened",
		)))
	})
})
