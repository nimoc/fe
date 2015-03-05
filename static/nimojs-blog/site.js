// http://fex-team.github.io/webuploader/js/global.js
// comments相关
$(function() {
    var $el = $('#ghComments'),
        $loading = $('#ghCommentsLoading'),
        issueId,
        ghname;

    // 不合法
    if ( !$el.length || !$el.attr('data-issue-id') ) {
        return;
    }

    issueId = $el.data('issue-id') ^ 0,
    ghname = $el.data('gh-name');

    function formatNumber(val, len) {
        var num = "" + val;

        len = len || 2;
        while (num.length < len) {
            num = "0" + num;
        }

        return num;
    }

    function formatDate( str ) {
        var date = new Date( str.replace(/T/, ' ').replace(/Z/, ' UTC') );

        return date.getFullYear() + '-' +
                formatNumber( date.getMonth() + 1 ) + '-' +
                formatNumber( date.getDate() ) + ' ' +
                formatNumber( date.getHours() ) + ':' +
                formatNumber( date.getMinutes() ) + ':' +
                formatNumber( date.getSeconds() );
    }

    function loadComments( data ) {
        $loading.hide();
        if (data.length === 0 ){
            $el.append('暂无评论');
        }
        for (var i = 0; i < data.length; i++) {
            var cuser = data[i].user.login;
            var cuserlink = 'https://www.github.com/' + data[i].user.login;
            var clink = 'https://github.com/' + ghname + '/issues/' + issueId + '#issuecomment-' + data[i].url.substring(data[i].url.lastIndexOf('/') + 1);
            var cbody = data[i].body_html;
            var cavatarlink = data[i].user.avatar_url;
            var cdate = formatDate( data[i].created_at );
            $el.append('<div class="comment"><div class="commentheader"><div class="commentgravatar"><img src="' + cavatarlink + '&s=40" alt="" width="20" height="20"></div><a class="commentuser" href="' + cuserlink + '" target="_blank">' + cuser + '</a><a class="commentdate" href="' + clink + '" target="_blank">' + cdate + '</a></div><div class="commentbody">' + cbody + '</div></div>');
        }
    }
    
    $.ajax('https://api.github.com/repos/' + ghname + '/issues/' + issueId + '/comments?per_page=100', {
        headers: {
            Accept: 'application/vnd.github.full+json'
        },
        dataType: 'json',
        success: loadComments
    });

    $("#emaillink").attr("href","mailto:nimo.jser@gmail.com");
});