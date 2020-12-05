#!/usr/bin/env bash

go run github.com/openconfig/ygot/generator -path=yang -output_file=aether_2_0_0/generated.go -package_name=aether_2_0_0 -generate_fakeroot \
       aether-subscriber@2020-10-22.yang apn-profile@2020-10-22.yang up-profile@2020-10-22.yang qos-profile@2020-10-22.yang \
       access-profile@2020-10-22.yang security-profile@2020-11-30.yang enterprise@2020-11-30.yang connectivity-service@2020-11-30.yang


sedi=(-i)
case "$(uname)" in
  Darwin*) sedi=(-i "")
esac

lf=$'\n'; sed "${sedi[@]}" -e "1s/^/\/\/ Code generated by YGOT. DO NOT EDIT.\\$lf/" aether_2_0_0/generated.go


