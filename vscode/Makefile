all: clean package

clean:
	@rm -rf dist/

package:
	@mkdir -p dist/
	@cd src/; vsce package --allow-missing-repository
	@cd src/; mv appetit*.vsix ../dist/