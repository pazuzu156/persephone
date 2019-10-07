require('./bootstrap');

require('bootstrap4-notify')
const hljs = require('highlight.js')

$.notifyDefaults({
    position: 'fixed',
    element: 'body',
    width: 'auto',
    placement: {
        from: 'top',
        align: 'center'
    },
    animate: {
		enter: 'animated fadeInDown',
		exit: 'animated fadeOutUp'
    },
    template: `<div data-notify="container" class="alert alert-{0} col-xs-11 col-sm-3" role="alert">
    <span data-notify="message">{2}</span>
</div>`
})

$(function() {
    $("#backToTopA").on('click', function (e) {
        e.preventDefault()
        $("html, body").animate({scrollTop: 0}, 800)
    })

    $(window).scroll(function () {
        var top = $("#backToTopA")

        if ($(this).scrollTop() > 300) {
            top.fadeIn()
        } else {
            top.fadeOut()
        }
    })

    $("pre code").each(function (i, block) {
        var classname = $(this).attr('class')
        $(this).removeClass(classname).addClass(classname.replace('language-', ''))
        hljs.highlightBlock(block)
    })
})
