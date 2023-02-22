data "illumio-core_vens" "all" {
    // optionally, this could be narrowed to specific VENs by using a query parameter
}

resource "null_resource" "ven-upgrade" {
  count = length(data.illumio-core_vens.all.items)
  provisioner "local-exec" {
      command = <<EOF
        curl -s -X PUT "${var.pce_url}/api/v2/orgs/${var.pce_org_id}/vens/upgrade" \
             -H 'Content-Type: application/json' \
             -u "${var.pce_api_key}:${var.pce_api_secret}" \
             -d '{"release": "${var.ven_version}", "vens": [{"href": "${data.illumio-core_vens.all.items[count.index].href}"}]}'
EOF
  }
}
