#version 120

void main(void) {
    gl_FragColor[0] = gl_FragCoord.x/468.0;
    gl_FragColor[1] = gl_FragCoord.y/320.0;
    gl_FragColor[2] = 1.0;
}
