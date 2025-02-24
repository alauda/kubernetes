/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/api/batch/v1"
	batchv1 "k8s.io/client-go/applyconfigurations/batch/v1"
	gentype "k8s.io/client-go/gentype"
	typedbatchv1 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

// fakeCronJobs implements CronJobInterface
type fakeCronJobs struct {
	*gentype.FakeClientWithListAndApply[*v1.CronJob, *v1.CronJobList, *batchv1.CronJobApplyConfiguration]
	Fake *FakeBatchV1
}

func newFakeCronJobs(fake *FakeBatchV1, namespace string) typedbatchv1.CronJobInterface {
	return &fakeCronJobs{
		gentype.NewFakeClientWithListAndApply[*v1.CronJob, *v1.CronJobList, *batchv1.CronJobApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("cronjobs"),
			v1.SchemeGroupVersion.WithKind("CronJob"),
			func() *v1.CronJob { return &v1.CronJob{} },
			func() *v1.CronJobList { return &v1.CronJobList{} },
			func(dst, src *v1.CronJobList) { dst.ListMeta = src.ListMeta },
			func(list *v1.CronJobList) []*v1.CronJob { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.CronJobList, items []*v1.CronJob) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
