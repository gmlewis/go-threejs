// Copyright 2016 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package three_test

import (
	"github.com/gmlewis/go-threejs/v2/three"
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
