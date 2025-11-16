{
  description = "Golang Template";

  inputs = { nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-25.05"; };

  outputs = { self, nixpkgs }:
    let
      # define system once
      system = "x86_64-linux";
      # use it here, and bind platform-specific packages to `pkgs`
      pkgs = nixpkgs.legacyPackages.${system};
      nativeBuildInputs = with pkgs; [ go ];
      developmentPakcages = with pkgs; [ gopls delve ];
    in {

      packages.${system}.default = pkgs.buildGoModule {
        pname = "golang google oauth tester" ;
        version = "1";
        src = ./.;
        vendorHash = "sha256-JFxW88GoE/oWL/JFcZZxL8AHQJr5O/cz1OoINzFS0Tk=";
        # subPackages = [ "path1", "path2" ]
        env = { CGO_ENABLED = 0; };
      };
      devShells.${system}.default = pkgs.mkShell {
        packages = nativeBuildInputs ++ developmentPakcages
          ++ [ pkgs.bashInteractive ];
      };
    };
}
