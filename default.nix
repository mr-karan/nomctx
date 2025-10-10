{ lib
, buildGoModule
, fetchFromGitHub
}:

buildGoModule rec {
  pname = "nomctx";
  version = "0.1.0";

  src = ./.;

  vendorHash = null;

  ldflags = [ "-s" "-w" ];

  meta = with lib; {
    description = "Fast context switching for Nomad";
    homepage = "https://github.com/karancode/nomctx";
    license = licenses.mit;
    maintainers = [ ];
    mainProgram = "nomctx";
  };
}
