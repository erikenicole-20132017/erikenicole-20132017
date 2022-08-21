class Course < ApplicationRecord
  has_many :enrollments, dependent: :destroy
  has_many :students, through: :enrollments, source: :user
  has_many :teachers, through: :enrollments, source: :user
  has_many :assignments, dependent: :destroy
  has_many :submissions, through: :assignments
  validates :name, presence: true
  validates :start_date, presence: true
  validates :end_date, presence: true
  validates :term, presence: true
  validates :year, presence: true
  def self.find_by_name(name)
    Course.find_by(name: name)
  end
end
