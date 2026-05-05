#!/usr/bin/env bash

bin_dir=$1/bin/

# Pinned by default to a recent controller-runtime release that fetches
# kubebuilder-tools archives from the kubernetes-sigs github releases
# rather than the deprecated kubebuilder-tools GCS bucket. Override with
# SETUP_ENVTEST_VERSION to track a different revision.
setup_envtest_version=${SETUP_ENVTEST_VERSION:-release-0.21}

mkdir -p ${bin_dir}
GOBIN=${bin_dir} go install "sigs.k8s.io/controller-runtime/tools/setup-envtest@${setup_envtest_version}"
