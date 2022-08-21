require 'gruff'
def draw_scatterplot(x_values, y_values)
  g = Gruff::Scatter.new(400)
  g.title = "GCD"
  g.x_axis_label = "X"
  g.y_axis_label = "Y"
  g.data('data', x_values, y_values)
  g.write("plot.png")
end
