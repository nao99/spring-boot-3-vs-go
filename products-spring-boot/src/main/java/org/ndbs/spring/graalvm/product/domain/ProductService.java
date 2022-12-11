package org.ndbs.spring.graalvm.product.domain;

import org.ndbs.spring.graalvm.product.api.model.CreateProductDto;
import org.ndbs.spring.graalvm.product.domain.model.Product;

import java.util.Set;

/**
 * ProductService interface <br>
 *
 * <br>
 *
 * @author Nikolai Osipov <nao99.dev@gmail.com>
 * @since  2022-12-11
 */
public interface ProductService {
    Set<Product> getProducts();
    Product getProductById(Long productId);
    Product createProduct(CreateProductDto createDto);
}
