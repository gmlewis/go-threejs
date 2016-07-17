if ( ! Detector.webgl ) {

	Detector.addGetWebGLMessage()
	document.getElementById( 'container' ).innerHTML = ""

}

var container, stats

var camera, controls, scene, renderer

var mesh, texture, geometry, material

var worldWidth = 128, worldDepth = 128,
worldHalfWidth = worldWidth / 2, worldHalfDepth = worldDepth / 2

var clock = new THREE.Clock()

init()
animate()

function init() {
	container = document.getElementById( 'container' )

	camera = new THREE.PerspectiveCamera( 60, window.innerWidth / window.innerHeight, 1, 20000 )
	camera.position.y = 200

	controls = new THREE.FirstPersonControls( camera )

	controls.movementSpeed = 500
	controls.lookSpeed = 0.1

	scene = new THREE.Scene()
	scene.fog = new THREE.FogExp2( 0xaaccff, 0.0007 )

	geometry = new THREE.PlaneGeometry( 20000, 20000, worldWidth - 1, worldDepth - 1 )
	geometry.rotateX( - Math.PI / 2 )

	for ( var i = 0, l = geometry.vertices.length; i < l; i ++ ) {
		geometry.vertices[ i ].y = 35 * Math.sin( i / 2 )
	}

	var texture = new THREE.TextureLoader().load( "http://threejs.org/examples/textures/water.jpg" )
	texture.wrapS = texture.wrapT = THREE.RepeatWrapping
	texture.repeat.set( 5, 5 )

	material = new THREE.MeshBasicMaterial( { color: 0x0044ff, map: texture } )

	mesh = new THREE.Mesh( geometry, material )
	scene.add( mesh )

	renderer = new THREE.WebGLRenderer()
	renderer.setClearColor( 0xaaccff )
	renderer.setPixelRatio( window.devicePixelRatio )
	renderer.setSize( window.innerWidth, window.innerHeight )

	container.innerHTML = ""

	container.appendChild( renderer.domElement )

	stats = new Stats()
	stats.domElement.style.position = 'absolute'
	stats.domElement.style.top = '0px'
	container.appendChild( stats.domElement )

	window.addEventListener( 'resize', onWindowResize, false )
}

function onWindowResize() {
	camera.aspect = window.innerWidth / window.innerHeight
	camera.updateProjectionMatrix()

	renderer.setSize( window.innerWidth, window.innerHeight )

	controls.handleResize()
}

function animate() {
	requestAnimationFrame( animate )

	render()
	stats.update()

}

function render() {
	var delta = clock.getDelta(),
		time = clock.getElapsedTime() * 10

	for ( var i = 0, l = geometry.vertices.length; i < l; i ++ ) {

		geometry.vertices[ i ].y = 35 * Math.sin( i / 5 + ( time + i ) / 7 )

	}

	mesh.geometry.verticesNeedUpdate = true

	controls.update( delta )
	renderer.render( scene, camera )
}

