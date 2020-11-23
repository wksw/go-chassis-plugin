package tenant_filte

import (
	"github.com/go-chassis/go-chassis/v2/core/loadbalancer"
	"github.com/go-chassis/go-chassis/v2/core/registry"
)

const (
	TenantFilte = "tenant_filte"
)

func init() {
	loadbalancer.InstallFilter(TenantFilte, FilterAvailableTenantAffinity)
}

// 优先选取上游租户内的实例
func FilterAvailableTenantAffinity(old []*registry.MicroServiceInstance, c []*loadbalancer.Criteria) []*registry.MicroServiceInstance {
	return nil
}
