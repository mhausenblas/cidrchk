package cidrchck

default contains = "no"
default overlaps = "no"

contains = "yes" {
    cr := input.targetcidr
    ior := input.incidr
    net.cidr_contains(cr, ior)
}

contains = "yes"  {
    cr := input.targetcidr
    ior := input.inip
    net.cidr_contains(cr, ior)
}

overlaps = "yes" {
    cr := input.targetcidr
    ior := input.incidr
    net.cidr_intersects(cr, ior)
}

expand[ips] {
    cr := input.incidr
    ips := net.cidr_expand(cr)
}