/**
 * First we will load all of this project's JavaScript dependencies which
 * includes Vue and other libraries. It is a great starting point when
 * building robust, powerful web applications using Vue and Laravel.
 */

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
        var scrollBottom = $(document).height() - $(window).height() - $(window).scrollTop()

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

// window.Vue = require('vue');

/**
 * The following block of code may be used to automatically register your
 * Vue components. It will recursively scan this directory for the Vue
 * components and automatically register them with their "basename".
 *
 * Eg. ./components/ExampleComponent.vue -> <example-component></example-component>
 */

// const files = require.context('./', true, /\.vue$/i);
// files.keys().map(key => Vue.component(key.split('/').pop().split('.')[0], files(key).default));

// Vue.component('example-component', require('./components/ExampleComponent.vue').default);

/**
 * Next, we will create a fresh Vue application instance and attach it to
 * the page. Then, you may begin adding components to this application
 * or customize the JavaScript scaffolding to fit your unique needs.
 */

// const app = new Vue({
//     el: '#app',
// });
