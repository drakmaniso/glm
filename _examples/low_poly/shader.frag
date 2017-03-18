#version 450 core

in PerVertex {
	layout(location = 0) flat in vec3 Normal;
  layout(location = 1) flat in uint Material;
} vertex;

out vec4 Color;

const vec4 palette[] = {
  {0.2, 0.0, 0.6, 1.0},
  {0.8, 0.2, 0.0, 1.0},
  {0.0, 0.3, 0.1, 1.0},
  {0.0, 0.2, 0.6, 1.0},
  {0.7, 0.0, 0.3, 1.0},
  {0.8, 0.5, 0.0, 1.0}
};

void main(void) {

  // Simplistic diffuse lighting
  const vec3 L = normalize(vec3(0.4, 0.6, 0.8));
  float NdotL = dot(vertex.Normal, L);
  float diff = clamp(NdotL, 0.2, 1.0);

	Color = diff * palette[vertex.Material];
}
