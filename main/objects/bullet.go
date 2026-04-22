components {
  id: "bullet"
  component: "/main/Scripts/bullet.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"bullet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/resources/weapons.atlas\"\n"
  "}\n"
  ""
}
