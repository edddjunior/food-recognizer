import os
import cv2
import tensorflow as tf
import numpy as np
import sys

def preprocess_image(img):
    resized_img = cv2.resize(img, (300, 300))
    normalized_img = resized_img.astype('float32') / 255.0
    return normalized_img

script_dir = os.path.dirname(os.path.abspath(__file__))
model_path = os.path.join(script_dir, 'saved_model')

model = tf.saved_model.load(model_path)

def predict_fruit(image_path):
    img = cv2.imread(image_path)
    if img is None:
        return "Error: Unable to load image."

    img = preprocess_image(img)
    img = np.expand_dims(img, axis=0)

    predictions = model(img)
    predicted_class_index = np.argmax(predictions, axis=1)
    fruit_labels = ['apple', 'banana']
    predicted_fruit = fruit_labels[predicted_class_index[0]]

    return predicted_fruit

if __name__ == "__main__":
    if len(sys.argv) != 2:
        sys.exit(1)

    image_path = sys.argv[1]
    result = predict_fruit(image_path)
    print(result)
