console.log("loaded");
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
    var earth = document.querySelector(".earth-layer");
    var galaxy = document.querySelector(".galaxy-layer");
    var sky = document.querySelector(".sky-layer");
    document.addEventListener('scroll',function() {
        var s = document.body.scrollTop || document.documentElement.scrollTop;
        earth.style.top=(280 + s*2)+"px";
        earth.style.webkitTransform = "scale("+1/(s/200+1)+")";
        earth.style.MozTransform = "scale("+1/(s/200+1)+")";
        earth.style.Transform = "scale("+1/(s/200+1)+")";
        galaxy.style.top=(s + s/10)+"px";
        if (s<0) s = 1;
        if (s==0) s = 1;

        sky.style.opacity=0.8/(s/500);
    })
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