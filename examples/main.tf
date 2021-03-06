terraform {
  required_version = ">= 0.12"
}

provider "appgate" {
  username = "admin"
  password = "admin"
  url      = "https://ec2-54-80-224-21.compute-1.amazonaws.com:444/admin"
  provider = "local"
  insecure = true
}


# data "appgate_appliance" "gateway_appliance" {
#   # appliance_id = "90eb0df9-14c0-45cc-bbd0-5dbd562b7d1b"
#   appliance_name = "gateway-0c8cbe4b-567d-4269-9143-7cccdd0f90ab-site1"
# }

data "appgate_site" "default_site" {
  site_name = "Default site"
}

resource "appgate_appliance" "new_gateway" {
  name     = "another-gateway"
  hostname = "envy-10-97-168-1337.devops"

  client_interface {
    hostname       = "envy-10-97-168-1338.devops"
    proxy_protocol = true
    https_port     = 447
    dtls_port      = 445
    allow_sources {
      address = "1.3.3.8"
      netmask = 0
      nic     = "eth0"
    }
    override_spa_mode = "UDP-TCP"
  }

  peer_interface {
    hostname   = "envy-10-97-168-1338.devops"
    https_port = "1338"

    allow_sources {
      address = "1.3.3.8"
      netmask = 0
      nic     = "eth0"
    }
  }


  admin_interface {
    hostname = "envy-10-97-168-1337.devops"
    https_ciphers = [
      "ECDHE-RSA-AES256-GCM-SHA384",
      "ECDHE-RSA-AES128-GCM-SHA256"
    ]
    allow_sources {
      address = "1.3.3.7"
      netmask = 0
      nic     = "eth0"
    }
  }

  tags = [
    "terraform",
    "api-created"
  ]
  notes = "hello world"
  site  = data.appgate_site.default_site.id

  connection {
    type     = "ssh"
    user     = "cz"
    password = "cz"
    host     = "10.97.168.30"
  }

  networking {

    hosts {
      hostname = "bla"
      address  = "0.0.0.0"
    }
    hosts {
      hostname = "foo"
      address  = "127.0.0.1"
    }

    nics {
      enabled = true
      name    = "eth0"

      ipv4 {
        dhcp {
          enabled = false
          dns     = true
          routers = true
          ntp     = true
        }

        static {
          address  = "10.10.10.1"
          netmask  = 24
          hostname = "appgate.company.com"
          snat     = true
        }

        static {
          address  = "20.20.20.1"
          netmask  = 32
          hostname = "test.company.com"
          snat     = false
        }
      }

      ipv6 {
        dhcp {
          enabled = true
          dns     = true
          ntp     = true
        }
        static {
          address  = "2001:db8:0:0:0:ff00:42:8329"
          netmask  = 24
          hostname = "appgate.company.com"
          snat     = true
        }

        static {
          address  = "2002:db8:0:0:0:ff00:42:1337"
          netmask  = 32
          hostname = "test.company.com"
          snat     = false
        }
      }

    }
    dns_servers = [
      "8.8.8.8",
      "1.1.1.1",
    ]
    dns_domains = [
      "aa.com"
    ]
    routes {
      address = "0.0.0.0"
      netmask = 24
      gateway = "1.2.3.4"
      nic     = "eth0"
    }
  }


  ntp {
    servers {
      hostname = "ntp.microsoft.com"
      key_type = "MD5"
      key      = "bla"
    }
    servers {
      hostname = "ntp.google.com"
      key_type = "MD5"
      key      = "bla"
    }
    # servers {
    #   hostname = "ntp.aws.com"
    #   key_type = "MD5"
    #   key      = "bla"
    # }
  }

  ssh_server {
    enabled                 = true
    port                    = 2222
    password_authentication = true
    allow_sources {
      address = "1.3.3.7"
      netmask = 0
      nic     = "eth0"
    }
    # allow_sources {
    #   address = "0.0.0.0"
    #   netmask = 0
    #   nic     = "eth1"
    # }
  }

  snmp_server {
    enabled    = false
    tcp_port   = 161
    udp_port   = 161
    snmpd_conf = "foo"
    allow_sources {
      address = "1.3.3.7"
      netmask = 0
      nic     = "eth0"
    }
  }

  healthcheck_server {
    enabled = true
    port    = 5555
    allow_sources {
      address = "1.3.3.7"
      netmask = 0
      nic     = "eth0"
    }
  }
  prometheus_exporter {
    enabled = true
    port    = 1234
    allow_sources {
      address = "1.3.3.7"
      netmask = 0
      nic     = "eth0"
    }
  }

  ping {
    allow_sources {
      address = "1.3.3.7"
      netmask = 0
      nic     = "eth0"
    }
  }

  # log_forwarder {
  #   enabled = true
  #   elasticsearch {
  #     url                      = "https://aws.com/elasticsearch/instance/asdaxllkmda64"
  #     aws_id                   = "string"
  #     aws_region               = "eu-west-2"
  #     use_instance_credentials = true
  #     retention_days           = 3
  #   }

  #   tcp_clients {
  #     name    = "Company SIEM"
  #     host    = "siem.company.com"
  #     port    = 8888
  #     format  = "json"
  #     use_tls = true
  #   }
  #   sites = [
  #     data.appgate_site.default_site.id
  #   ]
  # }

  # iot_connector {
  #   enabled = true
  #   clients {
  #     name      = "Printers"
  #     device_id = "12699e27-b584-464a-81ee-5b4784b6d425"

  #     sources {
  #       address = "1.3.3.7"
  #       netmask = 24
  #       nic     = "eth0"
  #     }
  #     snat = true
  #   }
  # }

  rsyslog_destinations {
    selector    = "*.*"
    template    = "%HOSTNAME% %msg%"
    destination = "10.10.10.2"
  }
  rsyslog_destinations {
    selector    = ":msg, contains, \"[AUDIT]\""
    template    = "%msg:9:$%"
    destination = "10.30.20.3"
  }

  hostname_aliases = [
    "appgatealias.company.com",
    "alias2.appgate.company.com"
  ]

  # https://sdphelp.appgate.com/adminguide/v5.1/about-appliances.html?anchor=controller-a
  # controller {
  #   enabled = true
  # }

  # https://sdphelp.appgate.com/adminguide/v5.1/about-appliances.html?anchor=logserver-a
  log_server {
    enabled = false
    # retention_days = 2
  }
  # https://sdphelp.appgate.com/adminguide/v5.1/about-appliances.html?anchor=gateway-a
  gateway {
    enabled = true
    vpn {
      weight = 60
      allow_destinations {
        address = "127.0.0.1"
        netmask = 0
        nic     = "eth0"
      }
    }
  }
  # Save the seed file locally in base 64 format.
  provisioner "local-exec" {
    command = "echo ${appgate_appliance.new_gateway.seed_file} > seed.b64"
  }
  # provisioner "remote-exec" {
  #   inline = [
  #     "echo ${appgate_appliance.new_gateway.seed_file}  > raw.b64",
  #     # "cat raw.b64 | base64 -d  | jq .  >> seed.json",
  #   ]
  # }

}

# output "seed_file" {
#   value = "${appgate_appliance.new_gateway.seed_file}"
# }

resource "appgate_ringfence_rule" "basic_rule" {
  name = "basic"
  tags = [
    "terraform",
    "api-created"
  ]

  actions {
    protocol  = "icmp"
    direction = "out"
    action    = "allow"

    hosts = [
      "10.10.20.0/24"
    ]

    types = [
      "0-255"
    ]

  }

  actions {
    protocol  = "tcp"
    direction = "in"
    action    = "allow"

    hosts = [
      "10.0.2.0/24"
    ]

    ports = [
      "22-25"
    ]
  }

}

resource "appgate_condition" "test_condition" {
  name = "teraform-example-condition"
  tags = [
    "terraform",
    "api-created"
  ]

  expression = <<-EOF
var result = false;
/*password*/
if (claims.user.hasPassword('test', 60)) {
  return true;
}
/*end password*/
return result;
EOF

  repeat_schedules = [
    "1h",
    "13:32"
  ]

  remedy_methods {
    type        = "DisplayMessage"
    message     = "This resoure requires you to enter your password again"
  }

}



resource "appgate_policy" "basic_policy" {
  name  = "terraform policy"
  notes = "terraform policy notes"
  tags = [
    "terraform",
    "api-created"
  ]
  disabled = false

  expression = <<-EOF
var result = false;
/*claims.user.groups*/
if(claims.user.groups && claims.user.groups.indexOf("developers") >= 0) {
  return true;
}
/*end claims.user.groups*/
/*criteriaScript*/
if (admins(claims)) {
  return true;
}
/*end criteriaScript*/
return result;
EOF
}


resource "appgate_criteria_script" "test_criteria_script" {
  name       = "adminOnly"
  expression = "return claims.user.username === 'admin';"
  notes      = "Only allow admin user"
  tags = [
    "terraform",
    "api-created"
  ]
}


resource "appgate_device_script" "example_device_script" {
  name     = "device_script_name"
  filename = "script.sh"
  content  = <<-EOF
#!/usr/bin/env bash
echo "hello world"
EOF
  tags = [
    "terraform",
    "api-created"
  ]
}

data "archive_file" "customization" {
  type        = "zip"
  output_path = "${path.module}/customization/package.zip"

  source {
    content  = <<-EOF
#!/usr/bin/env bash
echo "startup script"
EOF
    filename = "start"
  }

  source {
    content  = <<-EOF
#!/usr/bin/env bash
echo "stop script"
EOF
    filename = "stop"
  }
}

resource "appgate_appliance_customization" "test_customization" {
  name = "test customization"
  file = data.archive_file.customization.output_path

  tags = [
    "terraform",
    "api-created"
  ]
}


resource "appgate_ip_pool" "example_ip_pool" {
  name            = "ip range example"
  lease_time_days = 5
  ranges {
    first = "10.0.0.1"
    last  = "10.0.0.254"
  }

  tags = [
    "terraform",
    "api-created"
  ]
}
