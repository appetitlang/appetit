// Thanks to https://www.youtube.com/watch?v=q5V4T3o3CXE for a good chunk of this.
const vscode = require("vscode");

function activate(context) {
    let line_commenter = vscode.commands.registerCommand(
        'bryansmith.appetit.commentlines', function() {
            // Get the active editor
            const editor = vscode.window.activeTextEditor;
            
            // Get the edit builder
            editor.edit(editBuilder => {
                // For each selection
                for (const selection of editor.selections) {
                    // For each line in the selection
                    for (let i = selection.start.line; i <= selection.end.line; i++) {
                        // Get the beginning of the line
                        const line_start = new vscode.Position(i, 0);
                        // Add the comment symbol
                        editBuilder.insert(line_start, "- ")
                    }
                }
            })
        });

    let line_uncommenter = vscode.commands.registerCommand(
        'bryansmith.appetit.uncommentlines', function() {
            // Get the active editor
            const editor = vscode.window.activeTextEditor;
            // Get the edit builder
            editor.edit(editBuilder => {
                // For each selection
                for (const selection of editor.selections) {
                    // For each line in the selection
                    for (let i = selection.start.line; i <= selection.end.line; i++) {
                        const line = editor.document.lineAt(i)
                        const line_text = line.text
                        const range = new vscode.Range(line.range.start, line.range.end)
                        //const first_two_chars = line_text.substring(0, 1)

                        let uncommented_line = line_text.replace(/^-?\s*/, '');
                        uncommented_line = uncommented_line.trimStart();
                        //vscode.window.showInformationMessage("\"" + uncommented_line.trimStart() + "\"");
                        editBuilder.replace(range, uncommented_line)
                    }
                }
            })
        });
    
    context.subscriptions.push(line_commenter)
    context.subscriptions.push(line_uncommenter)
}

exports.activate = activate

module.exports = {
    activate
}