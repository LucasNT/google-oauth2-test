{
  description = "Teste do google oauth2";

  inputs = { nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-25.05"; };

  outputs = { self, nixpkgs }:
    let
      # define system once
      system = "x86_64-linux";
      # use it here, and bind platform-specific packages to `pkgs`
      pkgs = nixpkgs.legacyPackages.${system};
    in {
      devShells.${system}.default = pkgs.mkShell {
        packages = [ pkgs.go pkgs.gopls pkgs.bashInteractive pkgs.delve ];
      };
    };
}
