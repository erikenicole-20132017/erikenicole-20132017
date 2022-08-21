import static org.junit.Assert.*;
import org.junit.Test;
public class IsPrimeTest {
  // Math.isPrime(int) returns whether the given number is prime or not
  @Test
  public void testIsPrime() {
    assertTrue(Math.isPrime(2));
    assertTrue(Math.isPrime(3));
    assertTrue(Math.isPrime(5));
    assertTrue(Math.isPrime(7));
    assertTrue(Math.isPrime(11));
    assertTrue(Math.isPrime(13));
    assertTrue(Math.isPrime(17));
    assertTrue(Math.isPrime(19));
    assertTrue(Math.isPrime(23));
    assertTrue(Math.isPrime(29));
}
