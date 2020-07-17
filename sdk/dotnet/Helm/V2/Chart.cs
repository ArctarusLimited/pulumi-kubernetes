// Copyright 2016-2020, Pulumi Corporation.
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

// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

namespace Pulumi.Kubernetes.Helm.V2
{
    /// <summary>
    /// Chart is a component representing a collection of resources described by an arbitrary Helm
    /// Chart. The Chart can be fetched from any source that is accessible to the `helm` command
    /// line. Values in the `values.yml` file can be overridden using
    /// <see cref="BaseChartArgsUnwrap.Values" /> (equivalent to `--set` or having multiple
    /// `values.yml` files). Objects can be transformed arbitrarily by supplying callbacks to
    /// <see cref="BaseChartArgsUnwrap.Transformations" />.
    /// <para />
    /// <see cref="Chart"/> does not use Tiller. The Chart specified is copied and expanded locally;
    /// the semantics are equivalent to running `helm template` and then using Pulumi to manage the
    /// resulting YAML manifests. Any values that would be retrieved in-cluster are assigned fake
    /// values, and none of Tiller's server-side validity testing is executed.
    /// </summary>
    public sealed class Chart : ChartBase
    {
        /// <summary>
        /// Create a Chart resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Chart(string releaseName, Union<ChartArgs, LocalChartArgs> args, ComponentResourceOptions? options = null)
            : base(releaseName, args, options)
        {
        }
    }
}
