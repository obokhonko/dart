$(function() {
    $('a[href*=#]:not([href=#])').click(function() {

        $('.top-panel > a').removeClass('active');
        $(this).addClass('active');

        if (location.pathname.replace(/^\//,'') == this.pathname.replace(/^\//,'') && location.hostname == this.hostname) {
            var target = $(this.hash);
            target = target.length ? target : $('[name=' + this.hash.slice(1) +']');
            if (target.length) {
                $('html,body').animate({
                    scrollTop: target.offset().top - 60
                }, 1000);
                return false;
            }
        }
    });
    /*
    var earth = $('.earth-layer'),
        galaxy = $('.galaxy-layer'),
        sky = $('.sky-layer'),
        stars = $('.stars-layer'),
        window = $(window),
        height = window.height(),
        pos = 1;

    var backgroundShift = function() {
        var scale;
        console.log(height);
        sky.css({
            opacity: ((height * 2) - pos) / (height * 2) - 0.2
        });
        scale = Math.max(0, ((height * 0.5) - pos) / (height * 0.5));
        earth.css({'transform': 'translate3d(0,' + ((pos * 1.1) + height / 2) + 'px, 0)  scale3d(' + scale + ',' + scale + ',1 )'});
        galaxy.css({'transform': 'translate3d(0, ' + (pos * 1.05) + 'px, 0)'});
        stars.css({'transform': 'translate3d(0, ' + (pos * 1.02) + 'px, 0)'});
    };

    document.addEventListener('scroll', function() {
        backgroundShift();
    });
    */
});