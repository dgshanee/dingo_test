Dingo

##Project Overview
Dingo is a CLI tool used for building transportable command-line DOMs.

##Implementation Details
 - The DOM is rendered through a structure.dingo file, which contains all of the built structures.
 - The structures are then put into a map of what is rendered on each line.

 ##Components
  - Components are structs that are marshalled into JSON and put into structure.dingo.

##Display
 - How do we display the components? 

##Inheritance 
 - How does inheritance work?
 When a user creates an element, they can set the parent of that component to a predefined ID.
 The component is then stored in the "children" class of the parent.


 ##Render
  - How do we render the DOM tree?
  - There should be a body component that acts as the root node
  - From there, if a DOM is empty, a new body will be created and will act as the parent
