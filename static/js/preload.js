(function($) {
	function PreLoad(imgs, options) {
		this.imgs = (typeof imgs === 'string') ? [imgs] : imgs;
        this.opts = $.extend({}, PreLoad.DEFAULTS, options);

        this._unordered();
	}
	PreLoad.DEFAULTS = {
		each: null,
		all: null
	};
	PreLoad.prototype._unordered = function() {
		var imgs = this.imgs,
            opts = this.opts,
            count = 0,
            len = imgs.length;
        $.each(imgs, function(i, src) {
            if (typeof src != 'string') return;
            var imgObj = new Image();
            $(imgObj).on('load error', function() {
                opts.each && opts.each(count); //每张图片加载完后调用each函数
                if (count >= len - 1) {
                    opts.all && opts.all(); //所有图片加载完成后调用all函数
                }
                count++;
            });
            imgObj.src = src;
        });
	};
	$.extend({
		preload: function(imgs, opts) {
			new PreLoad(imgs, opts);
		}
	});
})(jQuery);
