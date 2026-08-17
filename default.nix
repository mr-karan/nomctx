{ lib
, buildGoModule
, fetchFromGitHub
}:

buildGoModule rec {
  pname = "nomctx";
  version = "0.1.0";

  src = ./.;

  vendorHash = "sha256-9XyQGqLcvBNC/EuweehHQheAcbTNqtdmu0KScNNfoLw=";

  ldflags = [ "-s" "-w" ];

  meta = with lib; {
    description = "Fast context switching for Nomad";
    homepage = "https://github.com/mr-karan/nomctx";
    license = licenses.mit;
    maintainers = [ ];
    mainProgram = "nomctx";
  };
}
