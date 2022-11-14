const touchHandler = function (element, communication) {
    const getCoordsAndPreventDefault = function (event) {
        event.preventDefault();
        const object = event.changedTouches[0];
        return [
            parseInt(object.clientX),
            parseInt(object.clientY)
        ];
    };
    let originalX, originalY, interX, interY, posX, posY;

    const touchStart = function (event) {
        [originalX, originalY] = [interX, interY] = [posX, posY] = getCoordsAndPreventDefault(event);
    };

    const touchMoved = function (event) {
        [posX, posY] = getCoordsAndPreventDefault(event);
        const x = -(interX - posX)
        const y = -(interY - posY)
        if (x === 0 && y === 0) {
            return;
        }
        communication.send({
            type: "mousemove",
            offset: {
                x: x,
                y: y
            }
        });

        interX = posX;
        interY = posY;
    };

    const touchEnd = function (event) {
        [posX, posY] = getCoordsAndPreventDefault(event);

        if (Math.abs(originalX - posX) < 5 && Math.abs(originalY - posY) < 5) {
            feedback();
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

};
