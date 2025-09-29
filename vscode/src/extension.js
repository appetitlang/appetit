// Thanks to https://www.youtube.com/watch?v=q5V4T3o3CXE for a good chunk of this.
const vscode = require("vscode");

function activate(context) {
    let line_commenter = vscode.commands.registerCommand(
        'bryansmith.appetit.commentlines', function() {
            // Get the actiev editor
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
    context.subscriptions.push(line_commenter)
}

exports.activate = activate

module.exports = {
    activate
}