import matplotlib.pyplot as plt
def draw_scatterplot(x_values, y_values):
    plt.scatter(x_values, y_values, s=20)
    plt.title("Scatter Plot")
    plt.xlabel("x values")
    plt.ylabel("y values")
    plt.show()
