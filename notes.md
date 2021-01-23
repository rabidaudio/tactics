TODO:

- error handling for initializing assets?
- animations, and limiting inputs while animations happen?


https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm

world cordinates, camera cordinates,
with units

Idea:
Animate from overhead perspective to side when battle is initiated.
Use map as plane but don't shear objects, so they appear to be sticking out
of plane.
Requires 3D animation equations, `ebiten` is only capable of 2D. the example
of perspective renders each line of an image at a different width scale, which
works for some angles but does not allow arbitrary camera motion.

`mgl32` has the math functions, but figuring out the arguments requires more
research. It may be too expensive for CPU.

- https://en.wikipedia.org/wiki/3D_projection#Mathematical_formula
- songho.ca/math/homogeneous/homogeneous.html
- https://towardsdatascience.com/how-to-transform-a-2d-image-into-a-3d-space-5fc2306e3d36
- https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#gaf73673a7e8e18ec6963e3774e6a94b87

Combat mechanics:

- Weapon Triangle (sword/lance/axe)
- Ranged advantage over mounted units? or mages? no flying units available

Melee: 1 tile range
Ranged: 2 tile range

Stats:

Atk, Def, Spd

- Swords: Def++, Spd--
- Axes: Atk++, Def--
- Lances: Spd++, Atk--

- Crossbow: Def++, Spd--
- Hunter: Atk++, Def--
- (ranged lance equivalent?)

specialty units:

- mounted lancer
- knight (sword) - lord?
- brute - (axe)

base stats?

- unit promotions
- improved weapon? or are weapons items that can be collected in-game?

magic:

- Wound - ranged attack
  - heavy dmg with cooldown? limited uses (per battle or per game)?
- Heal - melee attack against teammates for negative damage

experience? dmg dealt + bonus for kill?

# TODO

(incomplete list, in no particular order)

- [ ] combat mechanics
- [ ] ai
- [ ] balance game using ai
- [ ] combat animations
- [ ] leveling
- [ ] HUD
- [ ] character generation
- [ ] map generation (or just make maps)
- [ ] save/load
- [ ] story
- [ ] music
- [ ] sfx
- [ ] graphics


# Random ideas

- leave bodies on battlefield
- on return to same map in later plays, show gravestones of past play-through
- battle animation transition: switch from 2d to "3d" (map planes out, objects popup-book from map)
- unit promotion is like pokemon evolution animation
