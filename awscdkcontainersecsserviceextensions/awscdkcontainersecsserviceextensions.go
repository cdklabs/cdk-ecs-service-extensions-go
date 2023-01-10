package awscdkcontainersecsserviceextensions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.AliasedPortExtension",
		reflect.TypeOf((*AliasedPortExtension)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainerMutatingHook", GoMethod: "AddContainerMutatingHook"},
			_jsii_.MemberMethod{JsiiMethod: "addHooks", GoMethod: "AddHooks"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasPort", GoGetter: "AliasPort"},
			_jsii_.MemberProperty{JsiiProperty: "appProtocol", GoGetter: "AppProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "connectToService", GoMethod: "ConnectToService"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "containerMutatingHooks", GoGetter: "ContainerMutatingHooks"},
			_jsii_.MemberMethod{JsiiMethod: "modifyServiceProps", GoMethod: "ModifyServiceProps"},
			_jsii_.MemberMethod{JsiiMethod: "modifyTaskDefinitionProps", GoMethod: "ModifyTaskDefinitionProps"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "parentService", GoGetter: "ParentService"},
			_jsii_.MemberMethod{JsiiMethod: "prehook", GoMethod: "Prehook"},
			_jsii_.MemberMethod{JsiiMethod: "resolveContainerDependencies", GoMethod: "ResolveContainerDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberMethod{JsiiMethod: "useService", GoMethod: "UseService"},
			_jsii_.MemberMethod{JsiiMethod: "useTaskDefinition", GoMethod: "UseTaskDefinition"},
		},
		func() interface{} {
			j := jsiiProxy_AliasedPortExtension{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ServiceExtension)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.AliasedPortMutatingHook",
		reflect.TypeOf((*AliasedPortMutatingHook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "mutateContainerDefinition", GoMethod: "MutateContainerDefinition"},
		},
		func() interface{} {
			j := jsiiProxy_AliasedPortMutatingHook{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ContainerMutatingHook)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.AliasedPortMutatingHookProps",
		reflect.TypeOf((*AliasedPortMutatingHookProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.AliasedPortProps",
		reflect.TypeOf((*AliasedPortProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.AppMeshExtension",
		reflect.TypeOf((*AppMeshExtension)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainerMutatingHook", GoMethod: "AddContainerMutatingHook"},
			_jsii_.MemberMethod{JsiiMethod: "addHooks", GoMethod: "AddHooks"},
			_jsii_.MemberMethod{JsiiMethod: "connectToService", GoMethod: "ConnectToService"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "containerMutatingHooks", GoGetter: "ContainerMutatingHooks"},
			_jsii_.MemberMethod{JsiiMethod: "modifyServiceProps", GoMethod: "ModifyServiceProps"},
			_jsii_.MemberMethod{JsiiMethod: "modifyTaskDefinitionProps", GoMethod: "ModifyTaskDefinitionProps"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parentService", GoGetter: "ParentService"},
			_jsii_.MemberMethod{JsiiMethod: "prehook", GoMethod: "Prehook"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberMethod{JsiiMethod: "resolveContainerDependencies", GoMethod: "ResolveContainerDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "route", GoGetter: "Route"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberMethod{JsiiMethod: "useService", GoMethod: "UseService"},
			_jsii_.MemberMethod{JsiiMethod: "useTaskDefinition", GoMethod: "UseTaskDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "virtualNode", GoGetter: "VirtualNode"},
			_jsii_.MemberProperty{JsiiProperty: "virtualRouter", GoGetter: "VirtualRouter"},
			_jsii_.MemberProperty{JsiiProperty: "virtualService", GoGetter: "VirtualService"},
		},
		func() interface{} {
			j := jsiiProxy_AppMeshExtension{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ServiceExtension)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.AssignPublicIpDnsOptions",
		reflect.TypeOf((*AssignPublicIpDnsOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.AssignPublicIpExtension",
		reflect.TypeOf((*AssignPublicIpExtension)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainerMutatingHook", GoMethod: "AddContainerMutatingHook"},
			_jsii_.MemberMethod{JsiiMethod: "addHooks", GoMethod: "AddHooks"},
			_jsii_.MemberMethod{JsiiMethod: "connectToService", GoMethod: "ConnectToService"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "containerMutatingHooks", GoGetter: "ContainerMutatingHooks"},
			_jsii_.MemberProperty{JsiiProperty: "dns", GoGetter: "Dns"},
			_jsii_.MemberMethod{JsiiMethod: "modifyServiceProps", GoMethod: "ModifyServiceProps"},
			_jsii_.MemberMethod{JsiiMethod: "modifyTaskDefinitionProps", GoMethod: "ModifyTaskDefinitionProps"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parentService", GoGetter: "ParentService"},
			_jsii_.MemberMethod{JsiiMethod: "prehook", GoMethod: "Prehook"},
			_jsii_.MemberMethod{JsiiMethod: "resolveContainerDependencies", GoMethod: "ResolveContainerDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberMethod{JsiiMethod: "useService", GoMethod: "UseService"},
			_jsii_.MemberMethod{JsiiMethod: "useTaskDefinition", GoMethod: "UseTaskDefinition"},
		},
		func() interface{} {
			j := jsiiProxy_AssignPublicIpExtension{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ServiceExtension)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.AssignPublicIpExtensionOptions",
		reflect.TypeOf((*AssignPublicIpExtensionOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.AutoScalingOptions",
		reflect.TypeOf((*AutoScalingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.CloudwatchAgentExtension",
		reflect.TypeOf((*CloudwatchAgentExtension)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainerMutatingHook", GoMethod: "AddContainerMutatingHook"},
			_jsii_.MemberMethod{JsiiMethod: "addHooks", GoMethod: "AddHooks"},
			_jsii_.MemberMethod{JsiiMethod: "connectToService", GoMethod: "ConnectToService"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "containerMutatingHooks", GoGetter: "ContainerMutatingHooks"},
			_jsii_.MemberMethod{JsiiMethod: "modifyServiceProps", GoMethod: "ModifyServiceProps"},
			_jsii_.MemberMethod{JsiiMethod: "modifyTaskDefinitionProps", GoMethod: "ModifyTaskDefinitionProps"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parentService", GoGetter: "ParentService"},
			_jsii_.MemberMethod{JsiiMethod: "prehook", GoMethod: "Prehook"},
			_jsii_.MemberMethod{JsiiMethod: "resolveContainerDependencies", GoMethod: "ResolveContainerDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberMethod{JsiiMethod: "useService", GoMethod: "UseService"},
			_jsii_.MemberMethod{JsiiMethod: "useTaskDefinition", GoMethod: "UseTaskDefinition"},
		},
		func() interface{} {
			j := jsiiProxy_CloudwatchAgentExtension{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ServiceExtension)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.ConnectToProps",
		reflect.TypeOf((*ConnectToProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.Container",
		reflect.TypeOf((*Container)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainerMutatingHook", GoMethod: "AddContainerMutatingHook"},
			_jsii_.MemberMethod{JsiiMethod: "addHooks", GoMethod: "AddHooks"},
			_jsii_.MemberMethod{JsiiMethod: "connectToService", GoMethod: "ConnectToService"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "containerMutatingHooks", GoGetter: "ContainerMutatingHooks"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "modifyServiceProps", GoMethod: "ModifyServiceProps"},
			_jsii_.MemberMethod{JsiiMethod: "modifyTaskDefinitionProps", GoMethod: "ModifyTaskDefinitionProps"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parentService", GoGetter: "ParentService"},
			_jsii_.MemberMethod{JsiiMethod: "prehook", GoMethod: "Prehook"},
			_jsii_.MemberMethod{JsiiMethod: "resolveContainerDependencies", GoMethod: "ResolveContainerDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "trafficPort", GoGetter: "TrafficPort"},
			_jsii_.MemberMethod{JsiiMethod: "useService", GoMethod: "UseService"},
			_jsii_.MemberMethod{JsiiMethod: "useTaskDefinition", GoMethod: "UseTaskDefinition"},
		},
		func() interface{} {
			j := jsiiProxy_Container{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ServiceExtension)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.ContainerExtensionProps",
		reflect.TypeOf((*ContainerExtensionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.ContainerMutatingHook",
		reflect.TypeOf((*ContainerMutatingHook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "mutateContainerDefinition", GoMethod: "MutateContainerDefinition"},
		},
		func() interface{} {
			return &jsiiProxy_ContainerMutatingHook{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.CpuScalingProps",
		reflect.TypeOf((*CpuScalingProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.Environment",
		reflect.TypeOf((*Environment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultCloudMapNamespace", GoMethod: "AddDefaultCloudMapNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "capacityType", GoGetter: "CapacityType"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_Environment{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEnvironment)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.EnvironmentAttributes",
		reflect.TypeOf((*EnvironmentAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk-containers/ecs-service-extensions.EnvironmentCapacityType",
		reflect.TypeOf((*EnvironmentCapacityType)(nil)).Elem(),
		map[string]interface{}{
			"FARGATE": EnvironmentCapacityType_FARGATE,
			"EC2": EnvironmentCapacityType_EC2,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.EnvironmentProps",
		reflect.TypeOf((*EnvironmentProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.FireLensExtension",
		reflect.TypeOf((*FireLensExtension)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainerMutatingHook", GoMethod: "AddContainerMutatingHook"},
			_jsii_.MemberMethod{JsiiMethod: "addHooks", GoMethod: "AddHooks"},
			_jsii_.MemberMethod{JsiiMethod: "connectToService", GoMethod: "ConnectToService"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "containerMutatingHooks", GoGetter: "ContainerMutatingHooks"},
			_jsii_.MemberMethod{JsiiMethod: "modifyServiceProps", GoMethod: "ModifyServiceProps"},
			_jsii_.MemberMethod{JsiiMethod: "modifyTaskDefinitionProps", GoMethod: "ModifyTaskDefinitionProps"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parentService", GoGetter: "ParentService"},
			_jsii_.MemberMethod{JsiiMethod: "prehook", GoMethod: "Prehook"},
			_jsii_.MemberMethod{JsiiMethod: "resolveContainerDependencies", GoMethod: "ResolveContainerDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberMethod{JsiiMethod: "useService", GoMethod: "UseService"},
			_jsii_.MemberMethod{JsiiMethod: "useTaskDefinition", GoMethod: "UseTaskDefinition"},
		},
		func() interface{} {
			j := jsiiProxy_FireLensExtension{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ServiceExtension)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.FirelensMutatingHook",
		reflect.TypeOf((*FirelensMutatingHook)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "mutateContainerDefinition", GoMethod: "MutateContainerDefinition"},
		},
		func() interface{} {
			j := jsiiProxy_FirelensMutatingHook{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ContainerMutatingHook)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.FirelensProps",
		reflect.TypeOf((*FirelensProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.HttpLoadBalancerExtension",
		reflect.TypeOf((*HttpLoadBalancerExtension)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainerMutatingHook", GoMethod: "AddContainerMutatingHook"},
			_jsii_.MemberMethod{JsiiMethod: "addHooks", GoMethod: "AddHooks"},
			_jsii_.MemberMethod{JsiiMethod: "connectToService", GoMethod: "ConnectToService"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "containerMutatingHooks", GoGetter: "ContainerMutatingHooks"},
			_jsii_.MemberMethod{JsiiMethod: "modifyServiceProps", GoMethod: "ModifyServiceProps"},
			_jsii_.MemberMethod{JsiiMethod: "modifyTaskDefinitionProps", GoMethod: "ModifyTaskDefinitionProps"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parentService", GoGetter: "ParentService"},
			_jsii_.MemberMethod{JsiiMethod: "prehook", GoMethod: "Prehook"},
			_jsii_.MemberMethod{JsiiMethod: "resolveContainerDependencies", GoMethod: "ResolveContainerDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberMethod{JsiiMethod: "useService", GoMethod: "UseService"},
			_jsii_.MemberMethod{JsiiMethod: "useTaskDefinition", GoMethod: "UseTaskDefinition"},
		},
		func() interface{} {
			j := jsiiProxy_HttpLoadBalancerExtension{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ServiceExtension)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.HttpLoadBalancerProps",
		reflect.TypeOf((*HttpLoadBalancerProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@aws-cdk-containers/ecs-service-extensions.IEnvironment",
		reflect.TypeOf((*IEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultCloudMapNamespace", GoMethod: "AddDefaultCloudMapNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "capacityType", GoGetter: "CapacityType"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			return &jsiiProxy_IEnvironment{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk-containers/ecs-service-extensions.IGrantInjectable",
		reflect.TypeOf((*IGrantInjectable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "environmentVariables", GoMethod: "EnvironmentVariables"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
		},
		func() interface{} {
			j := jsiiProxy_IGrantInjectable{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IInjectable)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk-containers/ecs-service-extensions.IInjectable",
		reflect.TypeOf((*IInjectable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "environmentVariables", GoMethod: "EnvironmentVariables"},
		},
		func() interface{} {
			return &jsiiProxy_IInjectable{}
		},
	)
	_jsii_.RegisterInterface(
		"@aws-cdk-containers/ecs-service-extensions.ISubscribable",
		reflect.TypeOf((*ISubscribable)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "subscribe", GoMethod: "Subscribe"},
			_jsii_.MemberProperty{JsiiProperty: "subscriptionQueue", GoGetter: "SubscriptionQueue"},
		},
		func() interface{} {
			return &jsiiProxy_ISubscribable{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.ImportedEnvironment",
		reflect.TypeOf((*ImportedEnvironment)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDefaultCloudMapNamespace", GoMethod: "AddDefaultCloudMapNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "capacityType", GoGetter: "CapacityType"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_ImportedEnvironment{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IEnvironment)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.InjectableTopic",
		reflect.TypeOf((*InjectableTopic)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "environmentVariables", GoMethod: "EnvironmentVariables"},
			_jsii_.MemberMethod{JsiiMethod: "grant", GoMethod: "Grant"},
			_jsii_.MemberProperty{JsiiProperty: "topic", GoGetter: "Topic"},
		},
		func() interface{} {
			j := jsiiProxy_InjectableTopic{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IGrantInjectable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.InjectableTopicProps",
		reflect.TypeOf((*InjectableTopicProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.InjecterExtension",
		reflect.TypeOf((*InjecterExtension)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainerMutatingHook", GoMethod: "AddContainerMutatingHook"},
			_jsii_.MemberMethod{JsiiMethod: "addHooks", GoMethod: "AddHooks"},
			_jsii_.MemberMethod{JsiiMethod: "connectToService", GoMethod: "ConnectToService"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "containerMutatingHooks", GoGetter: "ContainerMutatingHooks"},
			_jsii_.MemberMethod{JsiiMethod: "modifyServiceProps", GoMethod: "ModifyServiceProps"},
			_jsii_.MemberMethod{JsiiMethod: "modifyTaskDefinitionProps", GoMethod: "ModifyTaskDefinitionProps"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parentService", GoGetter: "ParentService"},
			_jsii_.MemberMethod{JsiiMethod: "prehook", GoMethod: "Prehook"},
			_jsii_.MemberMethod{JsiiMethod: "resolveContainerDependencies", GoMethod: "ResolveContainerDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberMethod{JsiiMethod: "useService", GoMethod: "UseService"},
			_jsii_.MemberMethod{JsiiMethod: "useTaskDefinition", GoMethod: "UseTaskDefinition"},
		},
		func() interface{} {
			j := jsiiProxy_InjecterExtension{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ServiceExtension)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.InjecterExtensionProps",
		reflect.TypeOf((*InjecterExtensionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.MeshProps",
		reflect.TypeOf((*MeshProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk-containers/ecs-service-extensions.Protocol",
		reflect.TypeOf((*Protocol)(nil)).Elem(),
		map[string]interface{}{
			"HTTP": Protocol_HTTP,
			"TCP": Protocol_TCP,
			"HTTP2": Protocol_HTTP2,
			"GRPC": Protocol_GRPC,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.QueueAutoScalingOptions",
		reflect.TypeOf((*QueueAutoScalingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.QueueExtension",
		reflect.TypeOf((*QueueExtension)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainerMutatingHook", GoMethod: "AddContainerMutatingHook"},
			_jsii_.MemberMethod{JsiiMethod: "addHooks", GoMethod: "AddHooks"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingOptions", GoGetter: "AutoscalingOptions"},
			_jsii_.MemberMethod{JsiiMethod: "connectToService", GoMethod: "ConnectToService"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "containerMutatingHooks", GoGetter: "ContainerMutatingHooks"},
			_jsii_.MemberProperty{JsiiProperty: "eventsQueue", GoGetter: "EventsQueue"},
			_jsii_.MemberProperty{JsiiProperty: "logGroup", GoGetter: "LogGroup"},
			_jsii_.MemberMethod{JsiiMethod: "modifyServiceProps", GoMethod: "ModifyServiceProps"},
			_jsii_.MemberMethod{JsiiMethod: "modifyTaskDefinitionProps", GoMethod: "ModifyTaskDefinitionProps"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parentService", GoGetter: "ParentService"},
			_jsii_.MemberMethod{JsiiMethod: "prehook", GoMethod: "Prehook"},
			_jsii_.MemberMethod{JsiiMethod: "resolveContainerDependencies", GoMethod: "ResolveContainerDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberMethod{JsiiMethod: "useService", GoMethod: "UseService"},
			_jsii_.MemberMethod{JsiiMethod: "useTaskDefinition", GoMethod: "UseTaskDefinition"},
		},
		func() interface{} {
			j := jsiiProxy_QueueExtension{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ServiceExtension)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.QueueExtensionProps",
		reflect.TypeOf((*QueueExtensionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.ScaleOnCpuUtilization",
		reflect.TypeOf((*ScaleOnCpuUtilization)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainerMutatingHook", GoMethod: "AddContainerMutatingHook"},
			_jsii_.MemberMethod{JsiiMethod: "addHooks", GoMethod: "AddHooks"},
			_jsii_.MemberMethod{JsiiMethod: "connectToService", GoMethod: "ConnectToService"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "containerMutatingHooks", GoGetter: "ContainerMutatingHooks"},
			_jsii_.MemberProperty{JsiiProperty: "initialTaskCount", GoGetter: "InitialTaskCount"},
			_jsii_.MemberProperty{JsiiProperty: "maxTaskCount", GoGetter: "MaxTaskCount"},
			_jsii_.MemberProperty{JsiiProperty: "minTaskCount", GoGetter: "MinTaskCount"},
			_jsii_.MemberMethod{JsiiMethod: "modifyServiceProps", GoMethod: "ModifyServiceProps"},
			_jsii_.MemberMethod{JsiiMethod: "modifyTaskDefinitionProps", GoMethod: "ModifyTaskDefinitionProps"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parentService", GoGetter: "ParentService"},
			_jsii_.MemberMethod{JsiiMethod: "prehook", GoMethod: "Prehook"},
			_jsii_.MemberMethod{JsiiMethod: "resolveContainerDependencies", GoMethod: "ResolveContainerDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "scaleInCooldown", GoGetter: "ScaleInCooldown"},
			_jsii_.MemberProperty{JsiiProperty: "scaleOutCooldown", GoGetter: "ScaleOutCooldown"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberProperty{JsiiProperty: "targetCpuUtilization", GoGetter: "TargetCpuUtilization"},
			_jsii_.MemberMethod{JsiiMethod: "useService", GoMethod: "UseService"},
			_jsii_.MemberMethod{JsiiMethod: "useTaskDefinition", GoMethod: "UseTaskDefinition"},
		},
		func() interface{} {
			j := jsiiProxy_ScaleOnCpuUtilization{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ServiceExtension)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.Service",
		reflect.TypeOf((*Service)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addURL", GoMethod: "AddURL"},
			_jsii_.MemberProperty{JsiiProperty: "capacityType", GoGetter: "CapacityType"},
			_jsii_.MemberProperty{JsiiProperty: "cluster", GoGetter: "Cluster"},
			_jsii_.MemberMethod{JsiiMethod: "connectTo", GoMethod: "ConnectTo"},
			_jsii_.MemberProperty{JsiiProperty: "ecsService", GoGetter: "EcsService"},
			_jsii_.MemberMethod{JsiiMethod: "enableAutoScalingPolicy", GoMethod: "EnableAutoScalingPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "environment", GoGetter: "Environment"},
			_jsii_.MemberMethod{JsiiMethod: "getURL", GoMethod: "GetURL"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "scalableTaskCount", GoGetter: "ScalableTaskCount"},
			_jsii_.MemberProperty{JsiiProperty: "serviceDescription", GoGetter: "ServiceDescription"},
			_jsii_.MemberProperty{JsiiProperty: "targetGroup", GoGetter: "TargetGroup"},
			_jsii_.MemberProperty{JsiiProperty: "taskDefinition", GoGetter: "TaskDefinition"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vpc", GoGetter: "Vpc"},
		},
		func() interface{} {
			j := jsiiProxy_Service{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.ServiceBuild",
		reflect.TypeOf((*ServiceBuild)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.ServiceDescription",
		reflect.TypeOf((*ServiceDescription)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "add", GoMethod: "Add"},
			_jsii_.MemberProperty{JsiiProperty: "extensions", GoGetter: "Extensions"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
		},
		func() interface{} {
			return &jsiiProxy_ServiceDescription{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.ServiceExtension",
		reflect.TypeOf((*ServiceExtension)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainerMutatingHook", GoMethod: "AddContainerMutatingHook"},
			_jsii_.MemberMethod{JsiiMethod: "addHooks", GoMethod: "AddHooks"},
			_jsii_.MemberMethod{JsiiMethod: "connectToService", GoMethod: "ConnectToService"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "containerMutatingHooks", GoGetter: "ContainerMutatingHooks"},
			_jsii_.MemberMethod{JsiiMethod: "modifyServiceProps", GoMethod: "ModifyServiceProps"},
			_jsii_.MemberMethod{JsiiMethod: "modifyTaskDefinitionProps", GoMethod: "ModifyTaskDefinitionProps"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parentService", GoGetter: "ParentService"},
			_jsii_.MemberMethod{JsiiMethod: "prehook", GoMethod: "Prehook"},
			_jsii_.MemberMethod{JsiiMethod: "resolveContainerDependencies", GoMethod: "ResolveContainerDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberMethod{JsiiMethod: "useService", GoMethod: "UseService"},
			_jsii_.MemberMethod{JsiiMethod: "useTaskDefinition", GoMethod: "UseTaskDefinition"},
		},
		func() interface{} {
			return &jsiiProxy_ServiceExtension{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.ServiceProps",
		reflect.TypeOf((*ServiceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.SubscriptionQueue",
		reflect.TypeOf((*SubscriptionQueue)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.TopicSubscription",
		reflect.TypeOf((*TopicSubscription)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "queue", GoGetter: "Queue"},
			_jsii_.MemberMethod{JsiiMethod: "subscribe", GoMethod: "Subscribe"},
			_jsii_.MemberProperty{JsiiProperty: "subscriptionQueue", GoGetter: "SubscriptionQueue"},
			_jsii_.MemberProperty{JsiiProperty: "topic", GoGetter: "Topic"},
		},
		func() interface{} {
			j := jsiiProxy_TopicSubscription{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ISubscribable)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk-containers/ecs-service-extensions.TopicSubscriptionProps",
		reflect.TypeOf((*TopicSubscriptionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk-containers/ecs-service-extensions.XRayExtension",
		reflect.TypeOf((*XRayExtension)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addContainerMutatingHook", GoMethod: "AddContainerMutatingHook"},
			_jsii_.MemberMethod{JsiiMethod: "addHooks", GoMethod: "AddHooks"},
			_jsii_.MemberMethod{JsiiMethod: "connectToService", GoMethod: "ConnectToService"},
			_jsii_.MemberProperty{JsiiProperty: "container", GoGetter: "Container"},
			_jsii_.MemberProperty{JsiiProperty: "containerMutatingHooks", GoGetter: "ContainerMutatingHooks"},
			_jsii_.MemberMethod{JsiiMethod: "modifyServiceProps", GoMethod: "ModifyServiceProps"},
			_jsii_.MemberMethod{JsiiMethod: "modifyTaskDefinitionProps", GoMethod: "ModifyTaskDefinitionProps"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "parentService", GoGetter: "ParentService"},
			_jsii_.MemberMethod{JsiiMethod: "prehook", GoMethod: "Prehook"},
			_jsii_.MemberMethod{JsiiMethod: "resolveContainerDependencies", GoMethod: "ResolveContainerDependencies"},
			_jsii_.MemberProperty{JsiiProperty: "scope", GoGetter: "Scope"},
			_jsii_.MemberMethod{JsiiMethod: "useService", GoMethod: "UseService"},
			_jsii_.MemberMethod{JsiiMethod: "useTaskDefinition", GoMethod: "UseTaskDefinition"},
		},
		func() interface{} {
			j := jsiiProxy_XRayExtension{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ServiceExtension)
			return &j
		},
	)
}
