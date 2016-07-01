// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package three_test

import (
	"github.com/gmlewis/go-threejs/three"
)

func ExampleRaycaster() {
	/* t := */ three.New()

	/* TODO
	raycaster := t.NewRaycaster()
	mouse := t.NewVector2()

	   function onMouseMove( event ) {

	   	// calculate mouse position in normalized device coordinates
	   	// (-1 to +1) for both components

	   	mouse.x = ( event.clientX / window.innerWidth ) * 2 - 1;
	   	mouse.y = - ( event.clientY / window.innerHeight ) * 2 + 1;

	   }

	   function render() {

	   	// update the picking ray with the camera and mouse position
	   	raycaster.setFromCamera( mouse, camera );

	   	// calculate objects intersecting the picking ray
	   	var intersects = raycaster.intersectObjects( scene.children );

	   	for ( var i = 0; i < intersects.length; i++ ) {

	   		intersects[ i ].object.material.color.set( 0xff0000 );

	   	}

	   	renderer.render( scene, camera );

	   }

	   window.addEventListener( 'mousemove', onMouseMove, false );

	   window.requestAnimationFrame(render);
	*/
}
