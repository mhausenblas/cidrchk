package cidrchck

default contains = false
default overlaps = false

contains {
    cr := input.targetcidr
    ior := input.incidr
    net.cidr_contains(cr, ior)
}

contains {
    cr := input.targetcidr
    ior := input.inip
    net.cidr_contains(cr, ior)
}

overlaps {
    cr := input.targetcidr
    ior := input.incidr
    net.cidr_intersects(cr, ior)
}

expand[msg] {
    cr := input.incidr
    ips := net.cidr_expand(cr)
    msg := sprintf("%v", [ips])
}