var touchHandler = function (element, communication) {
    var originalX, originalY, interX, interY, posX, posY;

    var touchStart = function (event) {
        [originalX, originalY] = [interX, interY] = [posX, posY] = getCoordsAndPreventDefault(event);
    };

    var touchMoved = function (event) {
        [posX, posY] = getCoordsAndPreventDefault(event);
        communication.send({
          type: "mousemove",
          offset: {
            x: -(interX - posX),
            y: -(interY - posY)
          }
        });

        interX = posX;
        interY = posY;
    };

    var touchEnd = function (event) {
        [posX, posY] = getCoordsAndPreventDefault(event);

        if (Math.abs(originalX - posX) < 5 && Math.abs(originalY - posY) < 5) {
          communication.send({
            type: "mouseclick",
            commands: ["1"]
          });
        }

    };

    element.addEventListener('touchstart', touchStart.bind(this), false);
    element.addEventListener('touchmove', touchMoved.bind(this), false);
    element.addEventListener('touchend', touchEnd.bind(this), false);
    element.addEventListener('touchcancel', touchEnd.bind(this), false);

    var getCoordsAndPreventDefault = function (event) {
        event.preventDefault();
        var object = event.changedTouches[0];
        return [
            parseInt(object.clientX),
            parseInt(object.clientY)
        ];
    };
};
