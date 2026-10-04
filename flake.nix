{
  description = "Development tools for the portfolio";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";

  outputs = { nixpkgs, ... }:
    let
      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];
      forAllSystems = f: nixpkgs.lib.genAttrs systems (system: f nixpkgs.legacyPackages.${system});
    in
    {
      devShells = forAllSystems (pkgs: {
        default = pkgs.mkShell {
          packages = with pkgs; [
            go
            stdenv.cc
            gnumake
            air
            lsof
            tailwindcss_4
            (writeShellScriptBin "templ" ''
              exec ${go}/bin/go run github.com/a-h/templ/cmd/templ "$@"
            '')
          ];

          CGO_ENABLED = "1";
          GOFLAGS = "-mod=mod";
        };
      });
    };
}
