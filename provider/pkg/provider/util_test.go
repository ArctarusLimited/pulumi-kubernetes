// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"reflect"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/clientcmd"
	clientapi "k8s.io/client-go/tools/clientcmd/api"
)

func TestHasComputedValue(t *testing.T) {
	tests := []struct {
		name             string
		obj              *unstructured.Unstructured
		hasComputedValue bool
	}{
		{
			name:             "nil object does not have a computed value",
			obj:              nil,
			hasComputedValue: false,
		},
		{
			name:             "Empty object does not have a computed value",
			obj:              &unstructured.Unstructured{},
			hasComputedValue: false,
		},
		{
			name:             "Object with no computed values does not have a computed value",
			obj:              &unstructured.Unstructured{Object: map[string]interface{}{}},
			hasComputedValue: false,
		},
		{
			name: "Object with one concrete value does not have a computed value",
			obj: &unstructured.Unstructured{Object: map[string]interface{}{
				"field1": 1,
			}},
			hasComputedValue: false,
		},
		{
			name: "Object with one computed value does have a computed value",
			obj: &unstructured.Unstructured{Object: map[string]interface{}{
				"field1": 1,
				"field2": resource.Computed{},
			}},
			hasComputedValue: true,
		},
		{
			name: "Object with one nested computed value does have a computed value",
			obj: &unstructured.Unstructured{Object: map[string]interface{}{
				"field1": 1,
				"field2": map[string]interface{}{
					"field3": resource.Computed{},
				},
			}},
			hasComputedValue: true,
		},
		{
			name: "Object with nested maps and no computed values",
			obj: &unstructured.Unstructured{Object: map[string]interface{}{
				"field1": 1,
				"field2": map[string]interface{}{
					"field3": "3",
				},
			}},
			hasComputedValue: false,
		},
		{
			name: "Object with doubly nested maps and 1 computed value",
			obj: &unstructured.Unstructured{Object: map[string]interface{}{
				"field1": 1,
				"field2": map[string]interface{}{
					"field3": "3",
					"field4": map[string]interface{}{
						"field5": resource.Computed{},
					},
				},
			}},
			hasComputedValue: true,
		},
		{
			name: "Object with nested slice of map[string]interface{} has a computed value",
			obj: &unstructured.Unstructured{Object: map[string]interface{}{
				"field1": 1,
				"field2": []map[string]interface{}{
					{"field3": resource.Computed{}},
				},
			}},
			hasComputedValue: true,
		},
		{
			name: "Object with nested slice of interface{} has a computed value",
			obj: &unstructured.Unstructured{Object: map[string]interface{}{
				"field1": 1,
				"field2": []interface{}{
					resource.Computed{},
				},
			}},
			hasComputedValue: true,
		},
		{
			name: "Object with nested slice of map[string]interface{} with nested slice of interface{} has a computed value",
			obj: &unstructured.Unstructured{Object: map[string]interface{}{
				"field1": 1,
				"field2": []map[string]interface{}{
					{"field3": []interface{}{
						resource.Computed{},
					}},
				},
			}},
			hasComputedValue: true,
		},
		{
			name: "Complex nested object with computed value",
			obj: &unstructured.Unstructured{Object: map[string]interface{}{
				"field1": 1,
				"field2": []map[string]interface{}{
					{"field3": []interface{}{
						[]map[string]interface{}{
							{"field4": []interface{}{
								resource.Computed{},
							}},
						},
					}},
				},
			}},
			hasComputedValue: true,
		},
		{
			name: "Complex nested object with no computed value",
			obj: &unstructured.Unstructured{Object: map[string]interface{}{
				"field1": 1,
				"field2": []map[string]interface{}{
					{"field3": []interface{}{
						[]map[string]interface{}{
							{"field4": []interface{}{
								"field5",
							}},
						},
					}},
				},
			}},
			hasComputedValue: false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.hasComputedValue, hasComputedValue(test.obj), test.name)
	}
}

func TestFqName(t *testing.T) {
	obj := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "tests/v1alpha1",
			"kind":       "Test",
			"metadata": map[string]interface{}{
				"name": "myname",
			},
		},
	}

	if n := fqName(obj.GetNamespace(), obj.GetName()); n != "myname" {
		t.Errorf("Got %q for %v", n, obj)
	}

	obj.SetNamespace("mynamespace")
	if n := fqName(obj.GetNamespace(), obj.GetName()); n != "mynamespace/myname" {
		t.Errorf("Got %q for %v", n, obj)
	}
}

func Test_getActiveClusterFromConfig(t *testing.T) {
	const validKubeconfig = `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM1ekNDQWMrZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJeE1EUXlOekUxTkRjd05Wb1hEVE14TURReU5URTFORGN3TlZvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTmVoCkJNOUowSVkrZFI5UmZVSjI5SlRxYjF2U3QwZUsxNDN1aVBxZElJR3hiWFFvVmV6ZDhRUXloSUFsUG91Z0VWS0gKUjRoTFRreEZJS01XQ1F0dGNCdkVaRnZBRmtyeVBzWU81RWgxRjZHdzJNbDYvNWtvU1psM1hTMDVyN1hnTUdWTQp5cVJRaDMvVWJFcVZkWkRlRWlBSnh6N3JQSUMxc1FUSlVqVTZUY2JaRFVYVkdGMVZMck9uRkJlUmg1NkwwN2RiCjJTeGd3dFhmNVNTMEFlYnJrT0REYzUwUUdYc250UkZONzE5YnlhblZCc3VzWm5mZnZIRWs1bnE1NUFMdGE0bjcKNkZDR2pRNHhYY2hsYTVvMWlreityN2pMenJ5NlNsdHJQWU5ML2VYNHgvRU0xUFFuVktlUWloRTJoNzRyakhLcApibDRwNjZPSjhseWRGa0VKQWVNQ0F3RUFBYU5DTUVBd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0hRWURWUjBPQkJZRUZJWHlrY0tnZGI0SlhEM0tSelNKSG4rdlRCeXlNQTBHQ1NxR1NJYjMKRFFFQkN3VUFBNElCQVFBd1M1WjZvV056WnRhNE1EeWczNWJmcjVRaTlIdG5WN2c5Y2VOdmNUSStxd0d4VUhZUApnZzJSb1Q0ZU5MMkVQV1lBUmRrZVNtU3JYMGtFL1QydFRueGJFWENNdEI2TjhPRnZiQ3VnVzZPK1pwSDNwNHR0ClVFQ0UxT3ZiWHd5MkMvdmkyaXJuOWtEd3I3SkFVQ2FGRVppcmVPTWNDbGp6ZURNTDBDOUpqQlJOUmRqWHVscmIKSlRwL0RiWVJ0OFVJNW0zaVFIa2luakRHVkVhVHIzamVCTTZQakl1L25sakNlK1ovV0wyb3pFbzgydzN2cHpONQp2MGRvaHFkVmxPMzJnZDYrQlFORjhmUDI5bzBkT0NBalYvNHdCYmNjdUh1YnZCZ3U0cnFIc0hvZzM3MUVUdWwvCjlJbHFrZ2FmemVydVBzNms1UGFaUE1iK25nbzNZRG5ndkhuSAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
    server: https://kubernetes.docker.internal:6443
  name: docker-desktop
contexts:
- context:
    cluster: docker-desktop
    user: docker-desktop
  name: docker-desktop
current-context: docker-desktop
kind: Config
preferences: {}
users:
- name: docker-desktop
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGVENDQWYyZ0F3SUJBZ0lJWnBUZjVmbTZDQW93RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TVRBME1qY3hOVFEzTURWYUZ3MHlNakExTVRJeU16SXdNVEphTURZeApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1Sc3dHUVlEVlFRREV4SmtiMk5yWlhJdFptOXlMV1JsCmMydDBiM0F3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRRENMaDZ3MWV5WGpOQ3AKSzI5ZFJJQ3o3eHd3K1ZPVXVYYlh2R2NJTEFxaElUdlR3WTJqUmVaTFFXK3B5Wm9XUUdWZm5EYVZ5TGxmUUVaOQpXQm9IcEkvWGVvVWl4Uy9mWmVPN1RTeXA1bFpLcExzaXBMSE1RazN1NHp2d1RqelJITFJ4Q3k3b2RWTUVxVWFyCkkveUxiVUMxL1RkaGc3WkVZTFVrbEE4bWVhWFpHMGx5ZjA4UEdTdTVLUUJuTFVlbXk5OHNqV2U3YTBvdlRZd3kKTUhveUhyS0VGV0xCTmYrTm5TMTY3ZFBONzhTNCtENThobGxZTmZEZDVHbXJYYUFBYzVxeHhCSW5VcmhkSDJQawo2YUZkZXduQjFRQlV6OWVlVUJEVlFoQmwwbXNoMmRUSWF3cGlnbENrTnR1RlhoTEhGMitaRjFCSzN5VnZaYURsCmsyOTNnanlwQWdNQkFBR2pTREJHTUE0R0ExVWREd0VCL3dRRUF3SUZvREFUQmdOVkhTVUVEREFLQmdnckJnRUYKQlFjREFqQWZCZ05WSFNNRUdEQVdnQlNGOHBIQ29IVytDVnc5eWtjMGlSNS9yMHdjc2pBTkJna3Foa2lHOXcwQgpBUXNGQUFPQ0FRRUFnT1dxR2Q0TnlCRzFDOUovb1NVTmxzdkxSWXp4eEluZ0VsT09MUmlNN2t2dTduRXV6SHBYCkViODh3di9SSU1qWlFlbDFOTmdLWFJvb0hOSmpXcVppOG5aMEIxangySnNmaldrZWlPUE1aTjZqNzhzdDBqWmsKZDErSW5Oem1raEo4ck92UjVCd2xFcDNUcUtTN3J6dzF4MnphRkxUVWtZblh6Wnp4TkU3VGZuZVJVSG4zVyt4SwpXMFFaS3RkUlcvV0M5M3AvckcvZXp2Z0o5dCthenZwa2V1bklTUm5lbFpGQzgrZTR3ZXdoZm15TmRtUFVySThkCnQyMzhxeVhaNzZMTERKTFFDSTRieFRSVVpJM2NDdFY4bzU0UThnVHAvNklaRXVkV3dPbEdXa0FackdaMXNCN2QKQ1RRbjRVTVBXV0JmTzBPcFdZS3hGcVg4U1FpQndQaWhDdz09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBd2k0ZXNOWHNsNHpRcVN0dlhVU0FzKzhjTVBsVGxMbDIxN3huQ0N3S29TRTcwOEdOCm8wWG1TMEZ2cWNtYUZrQmxYNXcybGNpNVgwQkdmVmdhQjZTUDEzcUZJc1V2MzJYanUwMHNxZVpXU3FTN0lxU3gKekVKTjd1TTc4RTQ4MFJ5MGNRc3U2SFZUQktsR3F5UDhpMjFBdGYwM1lZTzJSR0MxSkpRUEpubWwyUnRKY245UApEeGtydVNrQVp5MUhwc3ZmTEkxbnUydEtMMDJNTWpCNk1oNnloQlZpd1RYL2paMHRldTNUemUvRXVQZytmSVpaCldEWHczZVJwcTEyZ0FIT2FzY1FTSjFLNFhSOWo1T21oWFhzSndkVUFWTS9YbmxBUTFVSVFaZEpySWRuVXlHc0sKWW9KUXBEYmJoVjRTeHhkdm1SZFFTdDhsYjJXZzVaTnZkNEk4cVFJREFRQUJBb0lCQVFDeitSY05BMW1EcFVvSQpZVytZWEZPRmNnc0pBUzJNWE5GZlp3bC9zNEl1a2FUbndTOUxzdytkbElxd0xXQ1pXeG9hSWFrZDdxcVJNL3VoClZUVGEvSlV0UEN1RmJJblFYcGxTRWxkaEtWRzFZVFRwQ1FpWnJxS1kxUmZLeEZqdDM5TUdLejFReXQwbEp0ZU8KNjQyNGxJd3pvUHZoYjdoUmEraTRmRm9HYVIxa09KN1dGcFNwM2pUa0pZckFpQWViL2IxUlZ5Rk9sNm9IcEozcQo3dmxoaHZibklJcXdrMHp4VU1ya1ArTDR2azhLdjhEcVZZMTg1M1B5UWJ3cm1EUnBkNWYvTmRwZ0lrMWNjUURZCk1OUUxPd3NaRThsdTJZMW9PcTVpRmhZZEFmM2o2Ykt2Vi92WFAxdXhtdlZSMEZ6eWU0L2JuaXBUcWdNYUI1ZnQKQWJ5MVJsL2hBb0dCQU5vRnBLVExmTGYvQUFqNGw4TmlvdjJ2OXhVWmlrQVhNL0NPREpDVzVEU2hZakNZV0tDNApYNm8vdlJ6bHF3NWNaMDEvKzBYZ2lRa0tBQUdacXlzRnRsYjgveDdkYThnVWVDRVhGcDUzNlRiZXdHaXJlSlRoCkNCSnhlQ0x1cjdLVmp6RnpHdFZlTzY4VzRHRjg0ZmlQaUk0ZVJETlExN0pYVVR2cGtZTGNCQ2RsQW9HQkFPUUIKU05hdS9GYVdHVG5DbkVKa2pqRm9JdWFGbnVLaHphd1FIUHVFSFV6TFMzT1Fqem54MVRmMGU5aWxkRHBoekJ0SQpoNUgrbzFvUmhNYlN5Z2g1SGQ5aE1nekM3cjcrNmdPU2hMOEdnNjJwNU13YzhSVUhnZWhOWmkxSEJaeUh0VGFFCmg3LzA2YjBOV3lyMDRVcGNSZXJIME44TWdSWXI2emZ4K25MblpGWDFBb0dBSE1kLzYwejlJcUNqbFl1VEpQU0IKUlhHVDhSSVZBTTdQU1dMRzM5TTdQb05MSGRVT1pmRFFsLzJmN2crWEcrY3dyN2RFS1A0eHVLSzhTM25JY1g1bwppbVVOSERyb1Bsb05YWGpad0lOZG9xT1d6SHBPQ1lFRytzQkZ0bjdCYkpaM2QzU1ZSek1RTjlXU091d3NQQTVlClhUdzdqbmFPY25rNlBPbGhEdUFTSUUwQ2dZRUFuNGpHam5DaDMzUG04cU5ZOHB1cFlxaWF3dkY3MnRlY01XaVUKM3VmeUdHbW13WlhFb2FhMHFoSkhGYSt2UTZwcVJpelpyeTJjM3NpalB2citvaThjMTlBS1ZTT1FLZFB6cWN3NwpWZTRZOU1xTGJNWlRhWU4zUWpQbDZvaG5STDh2N0pXTzVxRlhheENOV2VFK1FlbU9nbGlOcllQeVRyRXNSRmpzCkJMb2pXb0VDZ1lCaWMrWjJvSzNTTmpzL3J5ZFU1Lzg3T3NVbExHamxKbDI1NE0xaGl3RmVsd3pUWjNXWjFuZlkKcS80Mm5GR3VRQUQ3RFFwSTBCSnFWVTJCQlZySmlSeFhROVlXUStCb3Q5VU4yRVJQQmhFeityU0Y0MnhybnZobApsZTU2NHVmK3VBdCt2K2ZjZUtYVnVDRDN1ZGdxL2d5ejNCaHN5VkJxZkFoNy9oNndOTmhIb3c9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=
`
	// Outdented contexts[0].context.name
	const invalidKubeconfig = `apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM1ekNDQWMrZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRJeE1EUXlOekUxTkRjd05Wb1hEVE14TURReU5URTFORGN3TlZvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTmVoCkJNOUowSVkrZFI5UmZVSjI5SlRxYjF2U3QwZUsxNDN1aVBxZElJR3hiWFFvVmV6ZDhRUXloSUFsUG91Z0VWS0gKUjRoTFRreEZJS01XQ1F0dGNCdkVaRnZBRmtyeVBzWU81RWgxRjZHdzJNbDYvNWtvU1psM1hTMDVyN1hnTUdWTQp5cVJRaDMvVWJFcVZkWkRlRWlBSnh6N3JQSUMxc1FUSlVqVTZUY2JaRFVYVkdGMVZMck9uRkJlUmg1NkwwN2RiCjJTeGd3dFhmNVNTMEFlYnJrT0REYzUwUUdYc250UkZONzE5YnlhblZCc3VzWm5mZnZIRWs1bnE1NUFMdGE0bjcKNkZDR2pRNHhYY2hsYTVvMWlreityN2pMenJ5NlNsdHJQWU5ML2VYNHgvRU0xUFFuVktlUWloRTJoNzRyakhLcApibDRwNjZPSjhseWRGa0VKQWVNQ0F3RUFBYU5DTUVBd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0hRWURWUjBPQkJZRUZJWHlrY0tnZGI0SlhEM0tSelNKSG4rdlRCeXlNQTBHQ1NxR1NJYjMKRFFFQkN3VUFBNElCQVFBd1M1WjZvV056WnRhNE1EeWczNWJmcjVRaTlIdG5WN2c5Y2VOdmNUSStxd0d4VUhZUApnZzJSb1Q0ZU5MMkVQV1lBUmRrZVNtU3JYMGtFL1QydFRueGJFWENNdEI2TjhPRnZiQ3VnVzZPK1pwSDNwNHR0ClVFQ0UxT3ZiWHd5MkMvdmkyaXJuOWtEd3I3SkFVQ2FGRVppcmVPTWNDbGp6ZURNTDBDOUpqQlJOUmRqWHVscmIKSlRwL0RiWVJ0OFVJNW0zaVFIa2luakRHVkVhVHIzamVCTTZQakl1L25sakNlK1ovV0wyb3pFbzgydzN2cHpONQp2MGRvaHFkVmxPMzJnZDYrQlFORjhmUDI5bzBkT0NBalYvNHdCYmNjdUh1YnZCZ3U0cnFIc0hvZzM3MUVUdWwvCjlJbHFrZ2FmemVydVBzNms1UGFaUE1iK25nbzNZRG5ndkhuSAotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==
    server: https://kubernetes.docker.internal:6443
  name: docker-desktop
contexts:
- context:
    cluster: docker-desktop
    user: docker-desktop
name: docker-desktop
current-context: docker-desktop
kind: Config
preferences: {}
users:
- name: docker-desktop
  user:
    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURGVENDQWYyZ0F3SUJBZ0lJWnBUZjVmbTZDQW93RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWVGdzB5TVRBME1qY3hOVFEzTURWYUZ3MHlNakExTVRJeU16SXdNVEphTURZeApGekFWQmdOVkJBb1REbk41YzNSbGJUcHRZWE4wWlhKek1Sc3dHUVlEVlFRREV4SmtiMk5yWlhJdFptOXlMV1JsCmMydDBiM0F3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRRENMaDZ3MWV5WGpOQ3AKSzI5ZFJJQ3o3eHd3K1ZPVXVYYlh2R2NJTEFxaElUdlR3WTJqUmVaTFFXK3B5Wm9XUUdWZm5EYVZ5TGxmUUVaOQpXQm9IcEkvWGVvVWl4Uy9mWmVPN1RTeXA1bFpLcExzaXBMSE1RazN1NHp2d1RqelJITFJ4Q3k3b2RWTUVxVWFyCkkveUxiVUMxL1RkaGc3WkVZTFVrbEE4bWVhWFpHMGx5ZjA4UEdTdTVLUUJuTFVlbXk5OHNqV2U3YTBvdlRZd3kKTUhveUhyS0VGV0xCTmYrTm5TMTY3ZFBONzhTNCtENThobGxZTmZEZDVHbXJYYUFBYzVxeHhCSW5VcmhkSDJQawo2YUZkZXduQjFRQlV6OWVlVUJEVlFoQmwwbXNoMmRUSWF3cGlnbENrTnR1RlhoTEhGMitaRjFCSzN5VnZaYURsCmsyOTNnanlwQWdNQkFBR2pTREJHTUE0R0ExVWREd0VCL3dRRUF3SUZvREFUQmdOVkhTVUVEREFLQmdnckJnRUYKQlFjREFqQWZCZ05WSFNNRUdEQVdnQlNGOHBIQ29IVytDVnc5eWtjMGlSNS9yMHdjc2pBTkJna3Foa2lHOXcwQgpBUXNGQUFPQ0FRRUFnT1dxR2Q0TnlCRzFDOUovb1NVTmxzdkxSWXp4eEluZ0VsT09MUmlNN2t2dTduRXV6SHBYCkViODh3di9SSU1qWlFlbDFOTmdLWFJvb0hOSmpXcVppOG5aMEIxangySnNmaldrZWlPUE1aTjZqNzhzdDBqWmsKZDErSW5Oem1raEo4ck92UjVCd2xFcDNUcUtTN3J6dzF4MnphRkxUVWtZblh6Wnp4TkU3VGZuZVJVSG4zVyt4SwpXMFFaS3RkUlcvV0M5M3AvckcvZXp2Z0o5dCthenZwa2V1bklTUm5lbFpGQzgrZTR3ZXdoZm15TmRtUFVySThkCnQyMzhxeVhaNzZMTERKTFFDSTRieFRSVVpJM2NDdFY4bzU0UThnVHAvNklaRXVkV3dPbEdXa0FackdaMXNCN2QKQ1RRbjRVTVBXV0JmTzBPcFdZS3hGcVg4U1FpQndQaWhDdz09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBd2k0ZXNOWHNsNHpRcVN0dlhVU0FzKzhjTVBsVGxMbDIxN3huQ0N3S29TRTcwOEdOCm8wWG1TMEZ2cWNtYUZrQmxYNXcybGNpNVgwQkdmVmdhQjZTUDEzcUZJc1V2MzJYanUwMHNxZVpXU3FTN0lxU3gKekVKTjd1TTc4RTQ4MFJ5MGNRc3U2SFZUQktsR3F5UDhpMjFBdGYwM1lZTzJSR0MxSkpRUEpubWwyUnRKY245UApEeGtydVNrQVp5MUhwc3ZmTEkxbnUydEtMMDJNTWpCNk1oNnloQlZpd1RYL2paMHRldTNUemUvRXVQZytmSVpaCldEWHczZVJwcTEyZ0FIT2FzY1FTSjFLNFhSOWo1T21oWFhzSndkVUFWTS9YbmxBUTFVSVFaZEpySWRuVXlHc0sKWW9KUXBEYmJoVjRTeHhkdm1SZFFTdDhsYjJXZzVaTnZkNEk4cVFJREFRQUJBb0lCQVFDeitSY05BMW1EcFVvSQpZVytZWEZPRmNnc0pBUzJNWE5GZlp3bC9zNEl1a2FUbndTOUxzdytkbElxd0xXQ1pXeG9hSWFrZDdxcVJNL3VoClZUVGEvSlV0UEN1RmJJblFYcGxTRWxkaEtWRzFZVFRwQ1FpWnJxS1kxUmZLeEZqdDM5TUdLejFReXQwbEp0ZU8KNjQyNGxJd3pvUHZoYjdoUmEraTRmRm9HYVIxa09KN1dGcFNwM2pUa0pZckFpQWViL2IxUlZ5Rk9sNm9IcEozcQo3dmxoaHZibklJcXdrMHp4VU1ya1ArTDR2azhLdjhEcVZZMTg1M1B5UWJ3cm1EUnBkNWYvTmRwZ0lrMWNjUURZCk1OUUxPd3NaRThsdTJZMW9PcTVpRmhZZEFmM2o2Ykt2Vi92WFAxdXhtdlZSMEZ6eWU0L2JuaXBUcWdNYUI1ZnQKQWJ5MVJsL2hBb0dCQU5vRnBLVExmTGYvQUFqNGw4TmlvdjJ2OXhVWmlrQVhNL0NPREpDVzVEU2hZakNZV0tDNApYNm8vdlJ6bHF3NWNaMDEvKzBYZ2lRa0tBQUdacXlzRnRsYjgveDdkYThnVWVDRVhGcDUzNlRiZXdHaXJlSlRoCkNCSnhlQ0x1cjdLVmp6RnpHdFZlTzY4VzRHRjg0ZmlQaUk0ZVJETlExN0pYVVR2cGtZTGNCQ2RsQW9HQkFPUUIKU05hdS9GYVdHVG5DbkVKa2pqRm9JdWFGbnVLaHphd1FIUHVFSFV6TFMzT1Fqem54MVRmMGU5aWxkRHBoekJ0SQpoNUgrbzFvUmhNYlN5Z2g1SGQ5aE1nekM3cjcrNmdPU2hMOEdnNjJwNU13YzhSVUhnZWhOWmkxSEJaeUh0VGFFCmg3LzA2YjBOV3lyMDRVcGNSZXJIME44TWdSWXI2emZ4K25MblpGWDFBb0dBSE1kLzYwejlJcUNqbFl1VEpQU0IKUlhHVDhSSVZBTTdQU1dMRzM5TTdQb05MSGRVT1pmRFFsLzJmN2crWEcrY3dyN2RFS1A0eHVLSzhTM25JY1g1bwppbVVOSERyb1Bsb05YWGpad0lOZG9xT1d6SHBPQ1lFRytzQkZ0bjdCYkpaM2QzU1ZSek1RTjlXU091d3NQQTVlClhUdzdqbmFPY25rNlBPbGhEdUFTSUUwQ2dZRUFuNGpHam5DaDMzUG04cU5ZOHB1cFlxaWF3dkY3MnRlY01XaVUKM3VmeUdHbW13WlhFb2FhMHFoSkhGYSt2UTZwcVJpelpyeTJjM3NpalB2citvaThjMTlBS1ZTT1FLZFB6cWN3NwpWZTRZOU1xTGJNWlRhWU4zUWpQbDZvaG5STDh2N0pXTzVxRlhheENOV2VFK1FlbU9nbGlOcllQeVRyRXNSRmpzCkJMb2pXb0VDZ1lCaWMrWjJvSzNTTmpzL3J5ZFU1Lzg3T3NVbExHamxKbDI1NE0xaGl3RmVsd3pUWjNXWjFuZlkKcS80Mm5GR3VRQUQ3RFFwSTBCSnFWVTJCQlZySmlSeFhROVlXUStCb3Q5VU4yRVJQQmhFeityU0Y0MnhybnZobApsZTU2NHVmK3VBdCt2K2ZjZUtYVnVDRDN1ZGdxL2d5ejNCaHN5VkJxZkFoNy9oNndOTmhIb3c9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=
`
	certAuthData := []byte(`-----BEGIN CERTIFICATE-----
MIIC5zCCAc+gAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl
cm5ldGVzMB4XDTIxMDQyNzE1NDcwNVoXDTMxMDQyNTE1NDcwNVowFTETMBEGA1UE
AxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANeh
BM9J0IY+dR9RfUJ29JTqb1vSt0eK143uiPqdIIGxbXQoVezd8QQyhIAlPougEVKH
R4hLTkxFIKMWCQttcBvEZFvAFkryPsYO5Eh1F6Gw2Ml6/5koSZl3XS05r7XgMGVM
yqRQh3/UbEqVdZDeEiAJxz7rPIC1sQTJUjU6TcbZDUXVGF1VLrOnFBeRh56L07db
2SxgwtXf5SS0AebrkODDc50QGXsntRFN719byanVBsusZnffvHEk5nq55ALta4n7
6FCGjQ4xXchla5o1ikz+r7jLzry6SltrPYNL/eX4x/EM1PQnVKeQihE2h74rjHKp
bl4p66OJ8lydFkEJAeMCAwEAAaNCMEAwDgYDVR0PAQH/BAQDAgKkMA8GA1UdEwEB
/wQFMAMBAf8wHQYDVR0OBBYEFIXykcKgdb4JXD3KRzSJHn+vTByyMA0GCSqGSIb3
DQEBCwUAA4IBAQAwS5Z6oWNzZta4MDyg35bfr5Qi9HtnV7g9ceNvcTI+qwGxUHYP
gg2RoT4eNL2EPWYARdkeSmSrX0kE/T2tTnxbEXCMtB6N8OFvbCugW6O+ZpH3p4tt
UECE1OvbXwy2C/vi2irn9kDwr7JAUCaFEZireOMcCljzeDML0C9JjBRNRdjXulrb
JTp/DbYRt8UI5m3iQHkinjDGVEaTr3jeBM6PjIu/nljCe+Z/WL2ozEo82w3vpzN5
v0dohqdVlO32gd6+BQNF8fP29o0dOCAjV/4wBbccuHubvBgu4rqHsHog371ETul/
9IlqkgafzeruPs6k5PaZPMb+ngo3YDngvHnH
-----END CERTIFICATE-----
`)

	validConfig, _ := clientcmd.Load([]byte(validKubeconfig))
	outdentedConfig, _ := clientcmd.Load([]byte(invalidKubeconfig))

	type args struct {
		config    *clientapi.Config
		overrides resource.PropertyMap
	}
	tests := []struct {
		name string
		args args
		want *clientapi.Cluster
	}{
		{
			name: "nil",
			args: args{
				config:    nil,
				overrides: map[resource.PropertyKey]resource.PropertyValue{},
			},
			want: &clientapi.Cluster{},
		},
		{
			name: "valid",
			args: args{
				config:    validConfig,
				overrides: map[resource.PropertyKey]resource.PropertyValue{},
			},
			want: &clientapi.Cluster{
				Server:                   "https://kubernetes.docker.internal:6443",
				CertificateAuthorityData: certAuthData,
				Extensions:               map[string]runtime.Object{},
			},
		},
		{
			name: "invalid_context_override",
			args: args{
				config: validConfig,
				overrides: map[resource.PropertyKey]resource.PropertyValue{
					resource.PropertyKey("context"): {V: "foo"},
				},
			},
			want: &clientapi.Cluster{},
		},
		{
			name: "invalid_cluster_override",
			args: args{
				config: validConfig,
				overrides: map[resource.PropertyKey]resource.PropertyValue{
					resource.PropertyKey("cluster"): {V: "foo"},
				},
			},
			want: &clientapi.Cluster{},
		},
		{
			name: "outdented_context_name",
			args: args{
				config:    outdentedConfig,
				overrides: map[resource.PropertyKey]resource.PropertyValue{},
			},
			want: &clientapi.Cluster{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getActiveClusterFromConfig(tt.args.config, tt.args.overrides); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getActiveClusterFromConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
