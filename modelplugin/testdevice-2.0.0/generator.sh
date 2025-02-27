#!/usr/bin/env bash

go run github.com/openconfig/ygot/generator \
-path=../../../onos-helm-charts/config-models/testdevice-2.x/files/yang \
-output_file=testdevice_2_0_0/generated.go -package_name=testdevice_2_0_0 \
-generate_fakeroot onf-test1@2019-06-10.yang onf-test1-augmented@2020-02-29.yang

sedi=(-i)
case "$(uname)" in
  Darwin*) sedi=(-i "")
esac

lf=$'\n'; sed "${sedi[@]}" -e "1s/^/\/\/ Code generated by YGOT. DO NOT EDIT.\\$lf/" testdevice_2_0_0/generated.go


