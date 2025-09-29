# appetit README

This is a basic extension for Appetit for Visual Studio Code.

Features:
* Syntax highlighting support.
* Snippet support for each implemented statement, often accessible via the first two letters of each statement.
* Actions for Command Palette:
    * Commenting for selected text (Appetit: Comment Line(s))

### Snippets
`_` indicates a blank spot.
| Snippet Command | Output |
|----|----|
| ask | ask "" to "" |
| co | - _ |
| cd | copydirectory "" to "" |
| cp | copyfile "" to "" |
| dd | deletedirectory "" |
| df | deletefile "" |
| dl | download "" to "" |
| ex | execute "" |
| mkd | makedirectory "" |
| mkf | makefile "" |
| mi | minver _ |
| mvd | movedirectory "" to "" |
| mvf | movefile "" to "" |
| pa | pause _ |
| se | set _ = "" |
| wr | write "" |
| wrl | writeln "" |
| zd | zipdirectory "" to "" |
| zf | zipfile "" to "" |

Special thanks to:
- the official Batch File TextMate grammar for number support [here](https://github.com/microsoft/vscode/blob/main/extensions/bat/syntaxes/batchfile.tmLanguage.json).
- Olaf Neumann's Regex Generator [here](https://regex-generator.olafneumann.org/)

## Release Notes

### 1.0.0

Initial release of the extension. This is still in development so do not consider this stable yet.
