var path = require('path')
var fs = require('fs')
function tsdocHTML(content) {
    return tsdoc(content)
}
function tsdoc (content) {
    let md = content.trim().replace(/\;\`([\s\S]*?)^\`/gm, function(source, $1){
        return '```\r\n' + $1.replace(/\\\`/g,'`') + '```ts'
    })
    md = md.trim().replace(/^\`\`\`/,'')
    md +="\r\n```"
    md = md.trim()
    md = md.replace(/"\.\.\/lib\"/g,"\"percent-demo\"")

    /*
    支持语法:
    // @tsrun:hidden begin

    test('part 0  total 10 return 0', () = {
        expect(percent(0,10)).toBe(0)>
    })

    // @tsrun:hidden end
    */
    md = md.replace(/\/\/ \@tsrun:hidden begin[\s\S]*\/\/ \@tsrun:hidden end/g, '')

    md = md.replace(new RegExp("// \@ts-ignore.*", "g"), "")
    md = md.replace(/\`\`\`ts([\s\S]*)\`\`\`/gm, function (source, $1) {
        return '```ts\r\n' + $1.trim() + '\r\n```'
    })
    md = md.replace(/```ts;/, "```ts")
    md = md.replace(/\^\^\^/g, "```")
    md = md.replace(/\^/g, "`")
    md = md.trim()
    return md
}

fis.match("**", {
    release: false,
},1)
fis.match("(**).source.md", {
    parser: [
        function (content, file) {
            return content.replace(/\[(.*?)\|embed\]\((.*)\)/g, function (source, name, ref) {
                const extname = path.extname(ref)
                const code = fs.readFileSync(path.join(__dirname,file.subdirname, ref)).toString()
                return `[${name}](${ref})` +
                    '\r\n' +
                    '```' + extname +
                    '\r\n' +
                    code  +
                    '\r\n' +
                    '```'
            })
        }
    ],
    release: "/$1.md",
}, 999)

fis.match('(**).doc.ts', {
    parser: [
        function (content, file) {
            return tsdoc(content)
        }
    ],
    release: "/$1.md",
    isHtmlLike: true,
    rExt: "md",
}, 999)
