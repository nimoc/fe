var markrun = require('markrun')
fis.match('*.md', {
    rExt: '.html',
    isHtmlLike: true,
    parser: function (content, file) {
        var infoMarkrun = {
            filepath: file.fullname
        }
        var html = markrun(
            content,
            {},
            infoMarkrun
        )
        infoMarkrun.deps = infoMarkrun.deps || []
        infoMarkrun.deps.forEach(function (filename) {
             file.cache.addDeps(filename)
        })
    	html = html.replace(/href="([^"]+)\.md"/g, 'href="$1.html"')
        return html
    }
})
