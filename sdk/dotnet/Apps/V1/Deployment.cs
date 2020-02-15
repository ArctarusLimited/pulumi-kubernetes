// *** WARNING: this file was generated by the Pulumi Kubernetes codegen tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Apps.V1
{
    /// <summary>
    /// Deployment enables declarative updates for Pods and ReplicaSets.
    /// 
    /// This resource waits until its status is ready before registering success
    /// for create/update, and populating output properties from the current state of the resource.
    /// The following conditions are used to determine whether the resource creation has
    /// succeeded or failed:
    /// 
    /// 1. The Deployment has begun to be updated by the Deployment controller. If the current
    ///    generation of the Deployment is &amp;gt; 1, then this means that the current generation
    /// must
    ///    be different from the generation reported by the last outputs.
    /// 2. There exists a ReplicaSet whose revision is equal to the current revision of the
    ///    Deployment.
    /// 3. The Deployment's '.status.conditions' has a status of type 'Available' whose 'status'
    ///    member is set to 'True'.
    /// 4. If the Deployment has generation &amp;gt; 1, then '.status.conditions' has a status of
    /// type
    ///    'Progressing', whose 'status' member is set to 'True', and whose 'reason' is
    ///    'NewReplicaSetAvailable'. For generation &amp;lt;= 1, this status field does not exist,
    ///    because it doesn't do a rollout (i.e., it simply creates the Deployment and
    ///    corresponding ReplicaSet), and therefore there is no rollout to mark as 'Progressing'.
    /// 
    /// If the Deployment has not reached a Ready state after 10 minutes, it will
    /// time out and mark the resource update as Failed. You can override the default timeout value
    /// by setting the 'customTimeouts' option on the resource.
    /// </summary>
    public partial class Deployment : Pulumi.CustomResource
    {
        /// <summary>
        /// APIVersion defines the versioned schema of this representation of an object. Servers
        /// should convert recognized schemas to the latest internal value, and may reject
        /// unrecognized values. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
        /// </summary>
        [Output("apiVersion")]
        public Output<string> ApiVersion { get; private set; } = null!;

        /// <summary>
        /// Kind is a string value representing the REST resource this object represents. Servers
        /// may infer this from the endpoint the client submits requests to. Cannot be updated. In
        /// CamelCase. More info:
        /// https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Standard object metadata.
        /// </summary>
        [Output("metadata")]
        public Output<Types.Outputs.Meta.V1.ObjectMeta> Metadata { get; private set; } = null!;

        /// <summary>
        /// Specification of the desired behavior of the Deployment.
        /// </summary>
        [Output("spec")]
        public Output<Types.Outputs.Apps.V1.DeploymentSpec> Spec { get; private set; } = null!;

        /// <summary>
        /// Most recently observed status of the Deployment.
        /// </summary>
        [Output("status")]
        public Output<Types.Outputs.Apps.V1.DeploymentStatus> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Deployment resource with the given unique name, arguments, and options.
        /// </summary>
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Deployment(string name, Types.Inputs.Apps.V1.DeploymentArgs? args = null, CustomResourceOptions? options = null)
            : base("kubernetes:apps/v1:Deployment", name, SetAPIKindAndVersion(args), MakeResourceOptions(options))
        {
        }

        private static ResourceArgs SetAPIKindAndVersion(Types.Inputs.Apps.V1.DeploymentArgs? args)
        {
            args ??= new Types.Inputs.Apps.V1.DeploymentArgs();
            args.ApiVersion = "apps/v1";
            args.Kind = "Deployment";
            return args;
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            return CustomResourceOptions.Merge(defaultOptions, options);
        }

        /// <summary>
        /// Get an existing Deployment resource's state with the given name and ID.
        /// </summary>
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Deployment Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Deployment(name, null, CustomResourceOptions.Merge(options, new CustomResourceOptions
            {
                Id = id,
            }));
        }

    }
}
