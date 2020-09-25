package appgate

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccLdapIdentityProviderBasic(t *testing.T) {
	resourceName := "appgate_ldap_identity_provider.ldap_test_resource"
	rName := RandStringFromCharSet(10, CharSetAlphaNum)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLdapIdentityProviderDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLdapIdentityProviderBasic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckLdapIdentityProviderExists(resourceName),
					resource.TestCheckResourceAttr(resourceName, "admin_distinguished_name", "CN=admin,OU=Users,DC=company,DC=com"),
					resource.TestCheckResourceAttr(resourceName, "admin_password", "helloworld"),
					resource.TestCheckResourceAttr(resourceName, "admin_provider", "false"),
					resource.TestCheckResourceAttr(resourceName, "base_dn", "OU=Users,DC=company,DC=com"),
					resource.TestCheckResourceAttr(resourceName, "block_local_dns_requests", "true"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.#", "6"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.0.attribute_name", "objectGUID"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.0.claim_name", "userId"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.0.encrypted", "false"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.0.list", "false"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.1.attribute_name", "sAMAccountName"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.1.claim_name", "username"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.1.encrypted", "false"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.1.list", "false"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.2.attribute_name", "givenName"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.2.claim_name", "firstName"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.2.encrypted", "false"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.2.list", "false"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.3.attribute_name", "sn"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.3.claim_name", "lastName"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.3.encrypted", "false"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.3.list", "false"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.4.attribute_name", "mail"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.4.claim_name", "emails"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.4.encrypted", "false"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.4.list", "true"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.5.attribute_name", "memberOf"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.5.claim_name", "groups"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.5.encrypted", "false"),
					resource.TestCheckResourceAttr(resourceName, "claim_mappings.5.list", "true"),
					resource.TestCheckResourceAttr(resourceName, "default", "false"),
					resource.TestCheckResourceAttr(resourceName, "dns_search_domains.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dns_search_domains.0", "internal.company.com"),
					resource.TestCheckResourceAttr(resourceName, "dns_servers.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "dns_servers.0", "172.17.18.19"),
					resource.TestCheckResourceAttr(resourceName, "dns_servers.1", "192.100.111.31"),
					resource.TestCheckResourceAttr(resourceName, "hostnames.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "hostnames.0", "dc.ad.company.com"),
					resource.TestCheckResourceAttr(resourceName, "inactivity_timeout_minutes", "28"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_pool_v4"),
					resource.TestCheckResourceAttrSet(resourceName, "ip_pool_v6"),
					resource.TestCheckResourceAttr(resourceName, "membership_base_dn", "OU=Groups,DC=company,DC=com"),
					resource.TestCheckResourceAttr(resourceName, "membership_filter", "(objectCategory=group)"),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "notes", "Managed by terraform"),
					resource.TestCheckResourceAttr(resourceName, "object_class", "user"),
					resource.TestCheckResourceAttr(resourceName, "on_boarding_two_factor.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "on_demand_claim_mappings.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "password_warning.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "password_warning.0.enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "password_warning.0.message", "Your password is about to expire, Please change it"),
					resource.TestCheckResourceAttr(resourceName, "password_warning.0.threshold_days", "13"),
					resource.TestCheckResourceAttr(resourceName, "port", "389"),
					resource.TestCheckResourceAttr(resourceName, "ssl_enabled", "true"),
					resource.TestCheckResourceAttr(resourceName, "tags.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.2876187004", "api-created"),
					resource.TestCheckResourceAttr(resourceName, "tags.535570215", "terraform"),
					resource.TestCheckResourceAttr(resourceName, "type", "Ldap"),
					resource.TestCheckResourceAttr(resourceName, "username_attribute", "sAMAccountName"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateCheck:        testAccLdapIdentityProviderImportStateCheckFunc(1),
				ImportStateVerifyIgnore: []string{"admin_password"},
			},
		},
	})
}

func testAccCheckLdapIdentityProviderBasic(rName string) string {
	return fmt.Sprintf(`
data "appgate_ip_pool" "ip_four_pool" {
  ip_pool_name = "default pool v4"
}

data "appgate_ip_pool" "ip_sex_pool" {
  ip_pool_name = "default pool v6"
}

resource "appgate_ldap_identity_provider" "ldap_test_resource" {
  name                     = "%s"
  port                     = 389
  admin_distinguished_name = "CN=admin,OU=Users,DC=company,DC=com"
  hostnames                = ["dc.ad.company.com"]
  ssl_enabled              = true
  base_dn                  = "OU=Users,DC=company,DC=com"
  object_class             = "user"
  username_attribute       = "sAMAccountName"
  membership_filter        = "(objectCategory=group)"
  membership_base_dn       = "OU=Groups,DC=company,DC=com"
  password_warning {
    enabled        = true
    threshold_days = 13
    message        = "Your password is about to expire, Please change it"
  }
  default                    = false
  inactivity_timeout_minutes = 28
  ip_pool_v4                 = data.appgate_ip_pool.ip_four_pool.id
  ip_pool_v6                 = data.appgate_ip_pool.ip_sex_pool.id
  admin_password             = "helloworld"
  dns_servers = [
    "172.17.18.19",
    "192.100.111.31"
  ]
  dns_search_domains = [
    "internal.company.com"
  ]
  block_local_dns_requests = true
# TODO Update when we have mfa data source
#  on_boarding_two_factor {
#    mfa_provider_id       = "data.appgate_mfa_provider.id"
#    device_limit_per_user = 6
#    message               = "welcome"
#  }
  tags = [
    "terraform",
    "api-created"
  ]
}
`, rName)
}

func testAccCheckLdapIdentityProviderExists(resource string) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		token := testAccProvider.Meta().(*Client).Token
		api := testAccProvider.Meta().(*Client).API.LdapIdentityProvidersApi

		rs, ok := state.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Record ID is set")
		}

		_, _, err := api.IdentityProvidersIdGet(context.Background(), rs.Primary.ID).Authorization(token).Execute()
		if err != nil {
			return fmt.Errorf("error fetching ldap identity provider with resource %s. %s", resource, err)
		}
		return nil
	}
}

func testAccCheckLdapIdentityProviderDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "appgate_ldap_identity_provider" {
			continue
		}

		token := testAccProvider.Meta().(*Client).Token
		api := testAccProvider.Meta().(*Client).API.LdapIdentityProvidersApi

		_, _, err := api.IdentityProvidersIdGet(context.Background(), rs.Primary.ID).Authorization(token).Execute()
		if err == nil {
			return fmt.Errorf("ldap identity provider still exists, %+v", err)
		}
	}
	return nil
}

func testAccLdapIdentityProviderImportStateCheckFunc(expectedStates int) resource.ImportStateCheckFunc {
	return func(s []*terraform.InstanceState) error {
		if len(s) != expectedStates {
			return fmt.Errorf("expected %d states, got %d: %+v", expectedStates, len(s), s)
		}
		return nil
	}
}
